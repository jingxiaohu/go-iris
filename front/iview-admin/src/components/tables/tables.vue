<template>
  <div>
    <div v-if="searchable && searchPlace === 'top'" class="search-con search-con-top">
      <!--// FIXED 2 修复了searchKey是undefined的bug-->
      <Select v-model="searchKey" class="search-col">
        <Option v-for="item in columns" v-if="item.key !== 'handle' && item.type !== 'selection' && item.type !== 'index'" :value="item.key" :key="`search-col-${item.key}`">{{ item.title }}</Option>
      </Select>
      <Input clearable placeholder="输入关键字搜索" class="search-input" v-model="searchValue"
             @on-change="handleClear"
             @keydown.enter.native="handleSearch"/>
      <Button @click="handleSearch" class="search-btn" type="primary"><Icon type="ios-search"/>&nbsp;搜索</Button>
    </div>
    <Table
      ref="tablesMain"
      :data="insideTableData"
      :columns="insideColumns"
      :stripe="stripe"
      :border="border"
      :show-header="showHeader"
      :width="width"
      :height="height"
      :loading="loading"
      :disabled-hover="disabledHover"
      :highlight-row="highlightRow"
      :row-class-name="rowClassName"
      :size="size"
      :no-data-text="noDataText"
      :no-filtered-data-text="noFilteredDataText"
      @on-current-change="onCurrentChange"
      @on-select="onSelect"
      @on-select-cancel="onSelectCancel"
      @on-select-all="onSelectAll"
      @on-selection-change="onSelectionChange"
      @on-sort-change="onSortChange"
      @on-filter-change="onFilterChange"
      @on-row-click="onRowClick"
      @on-row-dblclick="onRowDblclick"
      @on-expand="onExpand"
    >
      <slot name="header" slot="header"></slot>
      <slot name="footer" slot="footer"></slot>
      <slot name="loading" slot="loading"></slot>
    </Table>
    <div style="margin: 10px;overflow: hidden">
      <Button v-if="showExport" type="primary" @click="exportExcel">导出为Csv文件</Button>
      <div v-if="showPage" style="float: right;">
        <Page :total="total" :current="1" size="small"
              @on-change="changePage" @on-page-size-change="changeSize"
              show-elevator show-sizer show-total/>
      </div>
    </div>

    <div v-if="searchable && searchPlace === 'bottom'" class="search-con search-con-top">
      <!--// FIXED 3 修复了searchKey是undefined的bug-->
      <Select v-model="searchKey" class="search-col">
        <Option v-for="item in columns" v-if="item.key !== 'handle' && item.type !== 'selection'" :value="item.key" :key="`search-col-${item.key}`">{{ item.title }}</Option>
      </Select>
      <Input clearable placeholder="输入关键字搜索" class="search-input" v-model="searchValue"
             @on-change="handleClear"
             @keydown.enter.native="handleSearch"/>
      <Button @click="handleSearch" class="search-btn" type="primary"><Icon type="ios-search"/>&nbsp;搜索</Button>
    </div>
    <a id="hrefToExportTable" style="display: none;width: 0px;height: 0px;"></a>
  </div>
</template>

<script>
import TablesEdit from './edit.vue'
import handleBtns from './handle-btns'
import './index.less'
export default {
  name: 'Tables',
  props: {
    showPage: {
      type: Boolean,
      default: true
    },
    showExport: {
      type: Boolean,
      default: true
    },
    total: Number,
    // 分页参数
    page: {
      type: Object,
      default: () => {
        return {
          start: 0,
          size: 10
        }
      }
    },

    value: {
      type: Array,
      default () {
        return []
      }
    },
    columns: {
      type: Array,
      default () {
        return []
      }
    },
    size: String,
    width: {
      type: [Number, String]
    },
    height: {
      type: [Number, String]
    },
    stripe: {
      type: Boolean,
      default: false
    },
    border: {
      type: Boolean,
      default: false
    },
    showHeader: {
      type: Boolean,
      default: true
    },
    highlightRow: {
      type: Boolean,
      default: false
    },
    rowClassName: {
      type: Function,
      default () {
        return ''
      }
    },
    context: {
      type: Object
    },
    noDataText: {
      type: String
    },
    noFilteredDataText: {
      type: String
    },
    disabledHover: {
      type: Boolean
    },
    loading: {
      type: Boolean,
      default: false
    },
    /**
     * @description 全局设置是否可编辑
     */
    editable: {
      type: Boolean,
      default: false
    },
    /**
     * @description 是否可搜索
     */
    searchable: {
      type: Boolean,
      default: false
    },
    /**
     * @description 搜索控件所在位置，'top' / 'bottom'
     */
    searchPlace: {
      type: String,
      default: 'top'
    }
  },
  /**
   * Events
   * @on-start-edit 返回值 {Object} ：同iview中render函数中的params对象 { row, index, column }
   * @on-cancel-edit 返回值 {Object} 同上
   * @on-save-edit 返回值 {Object} ：除上面三个参数外，还有一个value: 修改后的数据
   */
  data () {
    return {
      insideColumns: [],
      insideTableData: [],
      edittingCellId: '',
      edittingText: '',
      searchValue: '',
      searchKey: '',

      Page: {}
    }
  },
  methods: {
    suportEdit (item, index) {
      item.render = (h, params) => {
        return h(TablesEdit, {
          props: {
            params: params,
            value: this.insideTableData[params.index][params.column.key],
            edittingCellId: this.edittingCellId,
            editable: this.editable
          },
          on: {
            'input': val => {
              this.edittingText = val
            },
            'on-start-edit': (params) => {
              this.edittingCellId = `editting-${params.index}-${params.column.key}`
              this.$emit('on-start-edit', params)
            },
            'on-cancel-edit': (params) => {
              this.edittingCellId = ''
              this.$emit('on-cancel-edit', params)
            },
            'on-save-edit': (params) => {
              if (this.edittingText.trim().length !== 0) {
                this.value[params.row.initRowIndex][params.column.key] = this.edittingText
                this.$emit('input', this.value)
                this.$emit('on-save-edit', Object.assign(params, { value: this.edittingText }))
              }
              this.edittingCellId = ''
            }
          }
        })
      }
      return item
    },
    surportHandle (item) {
      let options = item.options || []
      let insideBtns = []
      options.forEach(item => {
        if (handleBtns[item]) insideBtns.push(handleBtns[item])
      })
      let btns = item.button ? [].concat(insideBtns, item.button) : insideBtns
      item.render = (h, params) => {
        params.tableData = this.value
        return h('div', btns.map(item => item(h, params, this)))
      }
      return item
    },
    handleColumns (columns) {
      this.insideColumns = columns.map((item, index) => {
        let res = item
        if (res.editable) res = this.suportEdit(res, index)
        if (res.key === 'handle') res = this.surportHandle(res)
        return res
      })
    },
    setDefaultSearchKey () {
      // FIXED 1 修复了searchKey是undefined的bug
      let filter = this.columns
        .filter(col => col.type !== 'selection')
        .filter(col => col.type !== 'index')
        .filter(col => col.key !== 'handle')
      if (filter.length > 0) {
        this.searchKey = filter[0].key
      }
      // this.searchKey = this.columns[0].key !== 'handle'
      //   ? this.columns[0].key
      //   : (this.columns.length > 1 ? this.columns[1].key : '')
    },
    handleClear (e) {
      if (e.target.value === '') this.insideTableData = this.value
    },
    handleSearch () {
      // FIXED 如果这里item的value是非字符串,会报错，所以转一下string
      this.insideTableData = this.value.filter(item => {
        // item[this.searchKey].indexOf(this.searchValue) > -1
        let val = item[this.searchKey]
        if (typeof (val) !== 'string') val = String(val)
        return val.indexOf(this.searchValue) > -1
      })
    },
    handleTableData () {
      this.insideTableData = this.value.map((item, index) => {
        let res = item
        res.initRowIndex = index
        return res
      })
    },
    exportCsv (params) {
      this.$refs.tablesMain.exportCsv(params)
    },
    clearCurrentRow () {
      this.$refs.talbesMain.clearCurrentRow()
    },
    onCurrentChange (currentRow, oldCurrentRow) {
      this.$emit('on-current-change', currentRow, oldCurrentRow)
    },
    onSelect (selection, row) {
      this.$emit('on-select', selection, row)
    },
    onSelectCancel (selection, row) {
      this.$emit('on-select-cancel', selection, row)
    },
    onSelectAll (selection) {
      this.$emit('on-select-all', selection)
    },
    onSelectionChange (selection) {
      this.$emit('on-selection-change', selection)
    },
    onSortChange (column, key, order) {
      this.$emit('on-sort-change', column, key, order)
    },
    onFilterChange (row) {
      this.$emit('on-filter-change', row)
    },
    onRowClick (row, index) {
      this.$emit('on-row-click', row, index)
    },
    onRowDblclick (row, index) {
      this.$emit('on-row-dblclick', row, index)
    },
    onExpand (row, status) {
      this.$emit('on-expand', row, status)
    },

    exportExcel () {
      this.$refs.tablesMain.exportCsv({
        filename: `table_${(new Date()).valueOf()}.csv`
      })
    },
    changePage (start) {
      // The simulated data is changed directly here, and the actual usage scenario should fetch the data from the server
      this.Page.start = start
      this.$emit('change-page', this.Page)
    },
    changeSize (size) {
      this.Page.size = size
      this.$emit('change-size', this.Page)
      // console.log(this.total, size)
    }
  },
  watch: {
    columns (columns) {
      this.handleColumns(columns)
      this.setDefaultSearchKey()
    },
    value (val) {
      this.handleTableData()
      if (this.searchable) this.handleSearch()
    }
  },
  mounted () {
    this.handleColumns(this.columns)
    this.setDefaultSearchKey()
    this.handleTableData()
    this.Page = this.page
  }
}
</script>
