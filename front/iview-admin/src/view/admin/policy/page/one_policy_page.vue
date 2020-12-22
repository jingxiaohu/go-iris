<template>
  <div>
    <Form ref="policyForm" :model="policy" :rules="ruleValidate" label-position="right" :label-width="110">
      <Row :gutter="2">
        <h4>权限信息</h4>
        <Col span="12" push="3">
        <FormItem label="租户：">
          <Select prefix="md-transgender" v-model="policy.v1" style="width: auto">
            <Option v-for="item in domainList" :value="item.value" :key="item.value">{{ item.label }}</Option>
          </Select>
        </FormItem>
        <FormItem prop="v5" label="权限名称：">
          <Input type="text" v-model="policy.v5" placeholder="请输入角色名称" clearable style="width: auto"/>
        </FormItem>
        <FormItem prop="v4" label="suf：">
          <Input type="text" v-model="policy.v4" placeholder="请输入suf" clearable style="width: auto"/>
        </FormItem>
        </Col>

        <Col span="12">
        <FormItem prop="v0" label="角色key：">
          <Select prefix="ios-infinite" v-model="policy.v0" style="width: 180px">
            <Option v-for="item in userGroupRoleList" :value="item.v1" :key="item.id">{{ item.v5 }}</Option>
          </Select>
          <!--<Input type="text" v-model="policy.v0" placeholder="请输入角色key" disabled style="width: auto"/>-->
        </FormItem>
        <FormItem prop="v2" label="URL策略：">
          <Input type="text" v-model="policy.v2" placeholder="请输入URL策略" clearable style="width: auto"/>
        </FormItem>
        <FormItem label="Action策略：">
          <Select multiple prefix="md-transgender" v-model="policy.v3" style="width: auto">
            <Option v-for="al in actionList" :value="al.value" :key="al.value">{{ al.label }}</Option>
          </Select>
        </FormItem>
        </Col>
      </Row>

      <h4>其他信息</h4>
      <br/>
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
    tPolicy: {
      type: Object,
      default: () => null
    },
    routeName: String
  },
  name: 'OnePolicyPage',
  data () {
    return {
      actionList: [
        { label: 'GET', value: 'GET' },
        { label: 'POST', value: 'POST' },
        { label: 'PUT', value: 'PUT' },
        { label: 'DELETE', value: 'DELETE' }
      ],
      domainList: [],
      policy: {v4: '.*'},
      ruleValidate: {
        v0: [
          { required: true, message: '请选择角色key', trigger: 'blur' }
        ],
        v1: [
          { required: true, message: '租户不能为空', trigger: 'blur' }
        ],
        v2: [
          { required: true, message: 'URL策略不能为空', trigger: 'blur' }
        ],
        v5: [
          { required: true, message: '角色名称不能为空', trigger: 'blur' }
        ]
      },
      userGroupRoleList: [],
    }
  },
  methods: {
    ...mapMutations([
      'closeTag'
    ]),
    ...mapActions([
      'handleUserGroupRoleList',
      'handleDomainList'
    ]),
    handleSubmit () {
      this.$refs.policyForm.validate(v => {
        if (v) {
          let { v1, v3 } = this.policy
          if (!v1) {
            this.$Message.error('租户不能为空')
            return
          }
          if (!v3 || v3.length === 0) {
            this.$Message.error('Action策略不能为空')
            return
          }
          //
          let p = { ...this.policy }
          if (v3.length === 1) p.v3 = v3.toString()
          if (v3.length > 1) p.v3 = v3.toString().replace(/,/g, '|')
          if (p.id) p.id = Number(p.id)
          this.$emit('submit', p)
        }
      })
    },

    handleReset () {
      if (this.tPolicy) {
        this.policy = { ...this.tPolicy }
        if (this.policy.v3) this.policy.v3 = this.policy.v3.split('|') || []
      }
      else this.policy = { v4: '.*' }
    },

    handleCloseTag () {
      let c = { name: this.routeName }
      if (this.tPolicy) c.query = this.tPolicy
      this.closeTag(c)
    },

    pullRoleList () {
      this.handleUserGroupRoleList().then(data => {
        if (data && data.length > 0) this.userGroupRoleList = data
      }).catch(err => console.log('错误：', err))
    },

    pullDomainList () {
      this.handleDomainList().then(data => {
        if (data) data.forEach(d => this.domainList.push({ value: d, label: d }))
      }).catch(err => console.log('错误：', err))
    },

  },
  created () {
    this.pullRoleList()
    this.pullDomainList()
    if (this.tPolicy) {
      this.policy = { ...this.tPolicy }
      if (this.policy.v3) this.policy.v3 = this.policy.v3.split('|') || []
    }
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
