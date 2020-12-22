<template>
  <div>
    <Row :gutter="3">
      <Col span="3">
        <Select v-model="params.domain" prefix="md-cloudy" placeholder="请选择租户" @on-change="selectChange">
          <Option v-for="dl in domainList" :value="dl.value" :key="dl.value">{{ dl.label }}</Option>
        </Select>
      </Col>
      <Col span="6">
        <Button class="tool" type="primary" icon="md-person-add" @click="$router.push({ path: '/add_role_page' })">新增</Button>
        <Button class="tool" type="warning" icon="md-create" @click="modifyRole" :disabled="modifyBtn">修改</Button>
        <Button class="tool" type="error" icon="md-close" @click="deleteRoles" :disabled="deleteBtn">删除</Button>
      </Col>
    </Row>
    <tables ref="tables"
            searchable
            search-place="top"
            v-model="tableData"
            :columns="columns"
            :showPage=false
            :highlightRow=true
            @on-delete="handleDelete"
            @on-selection-change="selectionChange"/>
  </div>
</template>

<script>
import Tables from '_c/tables'
import { mapActions } from 'vuex'
import neffos from 'neffos.js'
export default {
  components: {
    Tables
  },
  name: 'role',
  data () {
    return {
      modifyBtn: true,
      selection: null,
      //
      deleteBtn: true,
      deleteRoleKeys: [],

      domainList: [
        { value: 'all', label: '全部' }
      ],

      params: { domain: 'a' },
      tableData: [],
      columns: [
        { type: 'selection', align: 'center', width: 50 },
        { title: 'ID', key: 'id', sortable: true, width: 100 },
        // { title: '策略', key: 'pType', width: 80 },
        { title: '角色名称', key: 'v5', width: 200 },
        { title: '角色key', key: 'v1', width: 180 },
        // { title: '租户', key: 'v1', width: 85 },
        // { title: 'URL', key: 'v2', editable: true },
        // { title: 'action', key: 'v3', editable: true, width: 200 },
        // { title: 'suf', key: 'v4', editable: true, width: 80 },
        {
          title: '操作',
          key: 'handle',
          options: ['delete'],
        }
      ]
    }
  },
  methods: {
    ...mapActions([
      'handleDeleteRoles',
      'handleDomainList',
      'handleRoleTable'
    ]),
    handleDelete (params) {
      let id = params.row.v1
      this.deleteRoles(Array.of(id))
    },

    selectChange (params) {
      this.params.domain = params
      this.pullRoleTable()
    },

    modifyRole () {
      let tags = JSON.parse(localStorage.getItem('tagNaveList') || null)
      if (tags && tags.some(t => t.name === 'modify_role_page')) {
        this.$Message.warning('不能打开多个修改窗口')
        return
      }
      if (this.selection) this.$router.push({ name: 'modify_role_page', query: this.selection })
    },

    selectionChange (selection) {
      this.deleteRoleKeys = selection.map(a => a.v1);
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

    deleteRoles (ids) {
      let that = this
      this.$Modal.confirm({
        title: '删除角色',
        content: '确认删除选中的角色吗？',
        okText: '删除',
        cancelText: '取消',
        onOk () {
          that.handleDeleteRoles({ roleKeys: that.deleteRoleKeys.toString() }).then(res => {
            that.$Message.info({ content: '删除成功' })
            that.pullRoleTable(this.page)
            that.modifyBtn = true
            that.deleteBtn = true
          }).catch(err => console.log('删除错误：', err))
        }
      })
    },

    pullDomainList () {
      this.handleDomainList().then(data => {
        if (data) {
          data.forEach(d => {
            let dl = { value: d, label: d }
            this.domainList.push(dl)
          })
        }
      }).catch(err => console.log('错误：', err))
    },

    pullRoleTable () {
      this.handleRoleTable(this.params).then(data => {
        this.tableData = data || []
      }).catch(err => console.log('错误：', err))
    }
  },
  created () {
    this.pullDomainList()
    this.pullRoleTable()
  }
}
</script>

<style scoped>
  .tool {margin: 0 5px;}
</style>
