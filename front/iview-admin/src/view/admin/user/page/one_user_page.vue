<template>
  <div>
    <Form ref="userForm" :model="user" :rules="ruleValidate" label-position="right" :label-width="110">
      <Row :gutter="16">
        <h4>基本信息</h4>
        <Col span="12" push="2">
          <FormItem prop="name" label="用户姓名：">
            <Input type="text" v-model="user.name" placeholder="请输入用户姓名" clearable style="width: auto"/>
          </FormItem>
          <FormItem prop="phone" label="手机号码：">
            <Input type="text" v-model="user.phone" prefix="ios-phone-portrait" placeholder="请输入手机号码" clearable style="width: auto"/>
          </FormItem>
          <FormItem prop="username" label="登录账号：">
            <Input type="text" v-model="user.username" prefix="ios-contact" placeholder="请输入账号名称" clearable style="width: auto"/>
          </FormItem>
          <FormItem label="用户性别：">
            <Select prefix="md-transgender" v-model="user.gender" style="width: 194px">
              <Option v-for="item in genderList" :value="item.value" :key="item.value">{{ item.label }}</Option>
            </Select>
          </FormItem>
          <FormItem label="年龄：">
            <Input type="number" v-model="user.age" number placeholder="请输入年龄" style="width: auto"/>
          </FormItem>
          <FormItem label="岗位：">
            <Input type="text" placeholder="Enter text" clearable style="width: auto"/>
          </FormItem>
        </Col>

        <Col span="12" push="1">
          <FormItem label="所属角色：">
            <Select multiple prefix="ios-infinite" v-model="roles" style="width: auto">
              <Option v-for="item in userGroupRoleList" :value="item.v1" :key="item.id">{{ item.v5 }}</Option>
            </Select>
          </FormItem>
          <FormItem label="归属部门：">
            <Input type="text" number placeholder="Enter text" clearable style="width: auto"/>
          </FormItem>
          <FormItem label="邮箱：">
            <Input type="text" v-model="user.email" prefix="ios-mail" placeholder="请输入邮箱地址" clearable style="width: auto"/>
          </FormItem>
          <FormItem prop="password" label="登录密码：">
            <Input type="password" v-model="user.password" prefix="ios-lock" placeholder="请输入登录密码" clearable style="width: auto"/>
          </FormItem>
          <FormItem label="用户状态：">
            <i-switch size="large" v-model="user.enable">
              <span slot="open">ON</span>
              <span slot="close">OFF</span>
            </i-switch>
          </FormItem>
        </Col>
      </Row>

      <h4>其他信息</h4>
      <Row>
        <Col span="17" push="2">
          <FormItem label="备注：">
            <Input type="textarea" v-model="user.label" :rows="2" placeholder="描述"/>
          </FormItem>
        </Col>
      </Row>

      <FormItem :style="{padding: '0 400px'}">
        <Row>
          <Col span="8">
            <Button type="success" icon="md-checkmark" @click="handleSubmit">保存</Button>
          </Col>
          <Col span="8">
            <Button type="error" icon="md-return-left" @click="handleReset">重置</Button>
          </Col>
          <Col span="8">
            <Button type="warning" icon="md-close" @click="handleCloseTag">关闭</Button>
          </Col>
        </Row>
      </FormItem>

    </Form>
  </div>
</template>

<script>
import { mapMutations, mapActions } from 'vuex'
export default {
  props: {
    tUser: {
      type: Object,
      default: () => null
    },
    routeName: String
  },
  name: 'OneUserPage',
  data () {
    return {
      genderList: [
        { label: '男', 'value': 1 },
        { label: '女', 'value': 2 },
        { label: '未知', 'value': 0 }
      ],
      userGroupRoleList: [],
      roles: [],
      user: { enable: true },
      ruleValidate: {
        name: [
          { required: true, message: '名字不能为空', trigger: 'blur' }
        ],
        username: [
          { required: true, message: '登录账户不能为空', trigger: 'blur' }
        ],
        phone: [
          { required: true, message: '手机号不能为空', trigger: 'blur' }
        ],
        password: [
          { required: true, message: '登录密码不能为空', trigger: 'blur' }
        ],
        age: [
          { required: true, pattern: /^[a-z0-9]+$/, message: '请输入数字', trigger: 'change' }
        ],
        gender: [
          { required: true, message: '性别不能为空', trigger: 'change' }
        ]
      }
    }
  },
  methods: {
    ...mapMutations([
      'closeTag'
    ]),
    ...mapActions([
      'handleUserGroupRoleList'
    ]),

    findRole (v) {
      return this.userGroupRoleList.find(e => e.v1 === v)
    },

    handleSubmit () {
      this.$refs.userForm.validate(v => {
        if (v) {
          let result = { adminUser: { ...this.user } }
          let addRoles = []
          let deleteRids = []

          if (!this.tUser) { // 新建用户
            this.roles.forEach(e1 => {
              // let find = this.findRole(e1)
              addRoles.push({ v1: e1, v2: 'a', v5: this.findRole(e1).v5 })
            })
            result.addRoles = addRoles
          } else {
            let originRoles = this.user.roles || []
            // 添加的角色
            this.roles.forEach(e1 => {
              let find = originRoles.find(e2 => e1 === e2.v1)
              if (!find) addRoles.push({ v1: e1, v2: 'a', v5: this.findRole(e1).v5 })
            })
            // 删除的角色
            originRoles.forEach(e1 => {
              let find = this.roles.find(e2 => e1.v1 === e2)
              if (!find) deleteRids.push(e1.id)
            })
            result.addRoles = addRoles
            result.deleteRids = deleteRids
          }
          // console.log('原来拥有的角色id：', this.user.roles)
          // console.log('添加的角色：', addRoles)
          // console.log('删除的角色id：', deleteRids)
          // return
          this.$emit('submit', result)
        }
      })
    },

    handleReset () {
      if (this.tUser) this.user = { ...this.tUser }
      else this.user = { enable: true }
    },

    handleCloseTag () {
      this.closeTag({ name: this.routeName })
      sessionStorage.removeItem('selectionUser')
    },

    pullRoleList () {
      this.handleUserGroupRoleList().then(data => {
        if (data && data.length > 0) this.userGroupRoleList = data
      }).catch(err => console.log('错误：', err))
    }
  },
  created () {
    if (this.tUser) {
      this.user = { ...this.tUser }
      if (this.user.roles && this.user.roles.length > 0) this.roles = this.tUser.roles.map(e => e.v1)
    }
    this.pullRoleList()
  }
}
</script>
<style scoped>
  h4 {
    font-weight: 500;
    color: #6379bb;
    font-size: 15px;
    border-bottom: 1px solid #ddd;
    margin: 8px 10px 25px 10px;
    padding-bottom: 5px;
  }
</style>
