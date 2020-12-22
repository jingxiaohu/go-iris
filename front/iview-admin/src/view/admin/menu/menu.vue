<template>
  <div>
    <vxe-toolbar export custom>
      <template v-slot:buttons>
<!--        <vxe-button icon="vxe-icon&#45;&#45;plus" @click="$router.push({ path: '/add_menu_page' })">添加</vxe-button>-->
<!--        <vxe-button icon="vxe-icon&#45;&#45;edit-outline" @click="">修改</vxe-button>-->
<!--        <vxe-button icon="vxe-icon&#45;&#45;close" @click="">删除</vxe-button>-->
<!--        <vxe-button @click="expand">{{expandOption.text}}</vxe-button>-->
        <Button class="tool" type="primary" icon="md-person-add" @click="$router.push({ path: '/add_menu_page' })">新增</Button>
        <Button class="tool" type="warning" icon="md-create" @click="modify" :disabled="modifyBtn">修改</Button>
        <Button class="tool" type="error" icon="md-close" @click="deleteMenus" :disabled="deleteBtn">删除</Button>
        <Button class="tool" type="default" @click="expand">{{expandOption.text}}</Button>
      </template>
    </vxe-toolbar>
    <vxe-table
      border
      resizable
      row-key
      show-overflow
      highlight-hover-row
      highlight-current-row
      highlight-current-column
      ref="menuTree"
      :tree-config="{children: 'children', line: true}"
      :data="tableData"
      :checkbox-config="{labelField: 'id', highlight: true, checkStrictly: true}"
      @cell-click=""
      @select-change="selectionChange">
      <vxe-table-column type="checkbox" title="ID" sortable tree-node></vxe-table-column>
      <vxe-table-column field="meta" title="菜单名称" :formatter="fmtMeta"></vxe-table-column>
      <vxe-table-column field="parentId" title="父ID" width="70"></vxe-table-column>
      <vxe-table-column field="path" title="路径"></vxe-table-column>
      <vxe-table-column field="name" title="路径名称"></vxe-table-column>
      <vxe-table-column field="modular" title="模块"></vxe-table-column>
      <vxe-table-column field="component" title="组件"></vxe-table-column>
      <vxe-table-column field="sort" title="顺序" sortable width="70"></vxe-table-column>
      <vxe-table-column field="createTime" title="创建时间" sortable></vxe-table-column>
      <vxe-table-column field="updateTime" title="更新时间" sortable></vxe-table-column>
    </vxe-table>
    <vxe-pager
      :current-page="page.start"
      :page-size="page.size"
      :total="total"
      :layouts="['PrevPage', 'JumpNumber', 'NextPage', 'FullJump', 'Sizes', 'Total']"
      @page-change="handlePageChange">
    </vxe-pager>
  </div>
</template>

<script>
import { mapActions } from 'vuex'
export default {
  components: {
  },
  name: 'menu',
  data () {
    return {
      expandOption: {
        expand: false,
        text: '展开'
      },

      modifyBtn: true,
      selection: null,
      //
      deleteBtn: true,
      deleteIds: {mids: [], mMids: []},

      page: { start: 0, size: 10 },
      total: 0,
      tableData: []
    }
  },
  methods: {
    ...mapActions([
      'handleDeleteMenus',
      'handleMenuTable'
    ]),
    fmtMeta ({ cellValue, row, column }) {
      if (cellValue) return `${cellValue.title}`
      return ''
    },
    //
    expand () {
      if (this.expandOption.expand) {
        this.$refs.menuTree.clearTreeExpand()
        this.expandOption.text = '展开'
      } else {
        this.$refs.menuTree.setAllTreeExpansion(true)
        this.expandOption.text = '折叠'
      }
      this.expandOption.expand = !this.expandOption.expand
    },

    modify () {
      let tags = JSON.parse(localStorage.getItem('tagNaveList') || null)
      if (tags && tags.some(t => t.name === 'modify_menu_page')) {
        this.$Message.warning('不能打开多个修改窗口')
        return
      }
      // 复杂对象传递，用存本地存储
      if (this.selection) {
        sessionStorage.setItem('selectionMenu', JSON.stringify(this.selection))
        this.$router.replace({ name: 'modify_menu_page' })
      }
    },

    selectionChange ({ selection }) {
      this.deleteIds = {
        mids: selection.map(s => s.id),
        mMids: selection.map(s => s.meta.id)
      };
      (selection instanceof Array && selection.length === 1)
        ? this.modifyBtn = false
        : this.modifyBtn = true;
      (selection instanceof Array && selection.length > 0)
        ? this.deleteBtn = false
        : this.deleteBtn = true
      //
      if (selection.length === 1) this.selection = selection[0]
      else this.selection = null
    },

    deleteMenus () {
      let that = this
      this.$Modal.confirm({
        title: '删除菜单',
        content: '确认删除选中的菜单吗？',
        okText: '删除',
        cancelText: '取消',
        onOk () {
          that.handleDeleteMenus({ mids: that.deleteIds.mids.toString(), mMids: that.deleteIds.mMids.toString() }).then(res => {
            that.$Message.info({ content: '删除成功' })
            that.pullMenuTable()
            that.modifyBtn = true
            that.deleteBtn = true
          }).catch(err => console.log('删除错误：', err))
        }
      })
    },

    handlePageChange (p) {
      this.expandOption = {
        expand: false,
        text: '展开'
      }
      //
      this.page.start = p.currentPage
      this.page.size = p.pageSize
      this.pullMenuTable()
    },

    pullMenuTable () {
      this.handleMenuTable(this.page).then(data => {
        if (data) {
          this.total = data.total || 0
          this.tableData = data.data || []
        }
      }).catch(err => console.log('获取菜单报表数据错误：', err))
    }
  },
  created () {
    this.pullMenuTable()
  }
}

</script>
<style scoped>
  .tool {margin: 0 5px;}
</style>
