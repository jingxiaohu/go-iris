<template>
  <div>
    <Button class="tool" type="primary" icon="md-person-add" @click="$router.push({path: '/add_user_page'})">新增</Button>
    <Button class="tool" type="warning" icon="md-create" @click="modifyUser" :disabled="modifyBtn">修改</Button>
    <Button class="tool" type="error" icon="md-close" @click="deleteUsers(deleteIds)" :disabled="deleteBtn">删除</Button>
    <tables ref="tables"
            searchable
            highlightRow
            search-place="top"
            v-model="tableData"
            :columns="columns"
            :total="total"
            @change-page="pullUserTable"
            @change-size="pullUserTable"
            @on-delete="handleDelete"
            @on-selection-change="selectChange"/>
  </div>
</template>

<script>
import Tables from '_c/tables'
import { mapActions } from 'vuex'
export default {
  components: {
    Tables
  },
  name: 'user',
  data () {
    return {
      modifyBtn: true,
      selection: null,

      deleteBtn: true,
      deleteIds: [],

      page: { start: 0, size: 10 },
      total: 0,
      tableData: [],
      columns: [
        { type: 'selection', align: 'center', width: 35 },
        { title: 'ID', key: 'id', sortable: true, width: 70 },
        { title: '账号', key: 'username' },
        { title: '名称', key: 'name', width: 100 },
        { title: '角色', key: 'roles', width: 130,
          render: (h, params) => {
            let roles = params.row.roles
            if (roles && roles.length > 0) {
              let hh = []
              roles.forEach(r => hh.push(h('Tag', {props: {color: 'primary'}}, r.v5)))
              return h('div', hh)
            }
            return h('div', [
              h('Tag', {props: {color: 'default'}}, '无')
            ])
          }
        },
        { title: '年龄', key: 'age', sortable: true, width: 85 },
        { title: '启用', key: 'enable', width: 80,
          render: (h, params) => {
            let text = '未知'
            let color = 'default'
            let enable = params.row.enable
            switch (enable) {
              case true: text = '启用';color = 'success'; break;
              case false: text = '禁用';color = 'error';break;
              default: break;
            }
            return h('Tag', {
              props: { color: color }
            }, text)
          }
        },
        { title: '性别', key: 'gender', sortable: true, width: 100, align: 'center',
          render: (h, params) => {
            let text = '未知'
            let color = 'default'
            let gender = params.row.gender
            switch (gender) {
              case 1: text = '男';color = 'blue'; break;
              case 2: text = '女';color = 'green';break;
              default: break;
            }
            return h('Tag', {
              props: { color: color }
            }, text)
          }
        },
        { title: '电话', key: 'phone', editable: true },
        { title: '邮件', key: 'email', editable: true },
        { title: '头像', key: 'userface', editable: true },
        { title: '创建时间', key: 'createTime' },
        { title: '更新时间', key: 'updateTime' },
        { title: '备注', key: 'label' },
        {
          title: '操作',
          key: 'handle',
          options: ['delete']
        }
      ]
    }
  },
  methods: {
    ...mapActions([
      'handleDeleteUsers',
      'userTable'
    ]),
    handleDelete (params) {
      let id = params.row.id
      this.deleteUsers(Array.of(id))
    },

    modifyUser () {
      let tags = JSON.parse(localStorage.getItem('tagNaveList') || null)
      if (tags && tags.some(t => t.name === 'modify_user_page')) {
        this.$Message.warning('不能打开多个修改窗口')
        return
      }
      if (this.selection) {
        sessionStorage.setItem('selectionUser', JSON.stringify(this.selection))
        this.$router.replace({ name: 'modify_user_page' })
      }
      // if (this.selection) this.$router.push({ name: 'modify_user_page', query: this.selection })
    },

    selectChange (selection) {
      this.deleteIds = selection.map(a => a.id);
      (selection instanceof Array && selection.length === 1)
        ? this.modifyBtn = false
        : this.modifyBtn = true;
      (selection instanceof Array && selection.length > 0)
        ? this.deleteBtn = false
        : this.deleteBtn = true
      if (selection.length === 1) this.selection = selection[0]
      else this.selection = null
    },

    deleteUsers (ids) {
      let that = this
      this.$Modal.confirm({
        title: '删除用户',
        content: '确认删除选中的用户吗？',
        okText: '删除',
        cancelText: '取消',
        onOk () {
          that.handleDeleteUsers({ ids: ids.toString() }).then(res => {
            that.$Message.info({ content: '删除成功' })
            that.pullUserTable()
            that.modifyBtn = true
            that.deleteBtn = true
          }).catch(err => console.log('删除错误：', err))
        }
      })
    },

    pullUserTable () {
      this.userTable(this.page).then(data => {
        if (data) {
          this.total = data.total || 0
          this.tableData = data.data || []
        }
      }).catch(err => console.log('获取用户报表数据错误：', err))
    }
  },
  computed: {},
  mounted () {
    this.pullUserTable()
  }
}
</script>
<style scoped>
  .tool {margin: 0 5px;}
</style>
