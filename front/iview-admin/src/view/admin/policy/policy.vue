<template>
  <div>
    <Row :gutter="3">
      <Col span="3">
        <Select v-model="params.domain" prefix="md-cloudy" placeholder="请选择租户" @on-change="selectChange">
          <Option v-for="dl in domainList" :value="dl.value" :key="dl.value">{{ dl.label }}</Option>
        </Select>
      </Col>
      <Col span="6">
        <Button class="tool" type="primary" icon="md-person-add" @click="$router.push({ path: '/add_policy_page' })">新增</Button>
        <Button class="tool" type="warning" icon="md-create" @click="modifyRole" :disabled="modifyBtn">修改</Button>
        <Button class="tool" type="error" icon="md-close" @click="deleteRoles(deleteIds)" :disabled="deleteBtn">删除</Button>
      </Col>
    </Row>
    <tables ref="tables"
            searchable
            highlightRow
            search-place="top"
            v-model="tableData"
            :total="total"
            :columns="columns"
            @change-page="pullPolicyTable"
            @change-size="pullPolicyTable"
            @on-delete="handleDelete"
            @on-selection-change="selectionChange"/>
  </div>
</template>

<script>
import Tables from '_c/tables'
import { mapActions } from 'vuex'
export default {
  components: {
    Tables
  },
  name: 'policy',
  data () {
    return {
      modifyBtn: true,
      selection: null,
      //
      deleteBtn: true,
      deleteIds: [],

      domainList: [
        { value: 'all', label: '全部' }
      ],

      params: { start: 0, size: 10, domain: 'a' },
      total: 0,
      tableData: [],
      columns: [
        { type: 'selection', align: 'center', width: 50 },
        { title: 'ID', key: 'id', sortable: true, width: 80 },
        // { title: '策略', key: 'pType', width: 80 },
        { title: '权限名称', key: 'v5', width: 200 },
        { title: '关联的角色key', key: 'v0', width: 200 },
        { title: '租户', key: 'v1', width: 85 },
        { title: 'URL策略', key: 'v2', editable: true },
        { title: 'Action策略', key: 'v3', editable: true, width: 200 },
        { title: 'suf', key: 'v4', editable: true, width: 80 },
        { title: '操作', key: 'handle', options: ['delete'] }
      ]
    }
  },
  methods: {
    ...mapActions([
      'handleDeletePolicys',
      'handleDomainList',
      'handlePolicyTable'
    ]),
    handleDelete (params) {
      let id = params.row.id
      this.deleteRoles(Array.of(id))
    },
    selectChange (params) {
      this.params.domain = params
      this.pullPolicyTable()
    },
    modifyRole () {
      let tags = JSON.parse(localStorage.getItem('tagNaveList') || null)
      if (tags && tags.some(t => t.name === 'modify_policy_page')) {
        this.$Message.warning('不能打开多个修改窗口')
        return
      }
      if (this.selection) this.$router.push({ name: 'modify_policy_page', query: this.selection })
    },

    selectionChange (selection) {
      this.deleteIds = selection.map(a => a.id);
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
        title: '删除权限',
        content: '确认删除选中的权限规则吗？',
        okText: '删除',
        cancelText: '取消',
        onOk () {
          that.handleDeletePolicys({ ids: ids.toString() }).then(res => {
            that.$Message.info({ content: '删除成功' })
            that.pullPolicyTable()
            this.modifyBtn = true
            this.deleteBtn = true
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

    pullPolicyTable () {
      this.handlePolicyTable(this.params).then(data => {
        this.total = data.total || 0
        this.tableData = data.data || []
      }).catch(err => console.log('错误：', err))
    }
  },
  created () {
    this.pullDomainList()
    this.pullPolicyTable()
  }
}
</script>

<style scoped>
  .tool {margin: 0 5px;}
</style>
