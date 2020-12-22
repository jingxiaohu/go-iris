<template>
  <div>
    <Form ref="roleForm" :model="role" :rules="ruleValidate" label-position="right" :label-width="100">
      <Row :gutter="2">
        <h4>角色信息</h4>
        <Col span="7" push="1">
        <FormItem label="租户：">
          <Select prefix="md-transgender" v-model="role.v2" style="width: 170px">
            <Option v-for="item in domainList" :value="item.value" :key="item.value">{{ item.label }}</Option>
          </Select>
        </FormItem>
        </Col>
        <Col span="6">
        <FormItem prop="v5" label="角色名称：">
          <Input type="text" v-model="role.v5" placeholder="请输入角色名称" clearable style="width: auto"/>
        </FormItem>
        </Col>
        <Col span="6">
        <FormItem prop="v1" label="角色key：">
          <Input type="text" v-model="role.v1" prefix="ios-phone-portrait" placeholder="请输入角色key" clearable style="width: auto"/>
        </FormItem>
        </Col>
      </Row>

      <Row>
        <Col span="6" push="4">
        <Card title="关联菜单" :bordered="false">
          <Scroll>
            <Tree ref="menuTree" :data="menuList" show-checkbox />
          </Scroll>
        </Card>
        </Col>
        <Col span="6" push="6">
        <Card dis-hover :style="{background:'rgba(200,200,200,.2)'}" :bordered="false">
          <Tooltip slot="title" max-width="350">
            <p><Icon type="ios-information-circle-outline" size="18" />关联权限</p>
            <div slot="content">
              <h3>请在权限管理中设置角色的权限</h3>
              <h6>由于casbin原生的权限模型设计，这里不增加或删除权限</h6>
            </div>
          </Tooltip>
          <Scroll>
            <Tree ref="policyTree" :data="policyList" show-checkbox />
          </Scroll>
        </Card>
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
            <Button type="warning" icon="md-close" @click="handleCloseTag">关闭</Button>
          </Col>
        </Row>
      </FormItem>
    </Form>
  </div>
</template>

<script>
import { mapMutations, mapActions } from 'vuex'
import { Tooltip } from 'view-design'
export default {
  props: {
    tRole: {
      type: Object,
      default: () => null
    },
    routeName: String,
    isModify: Boolean
  },
  components: {
    Tooltip
  },
  name: 'OneRolePage',
  data () {
    return {
      actionList: [
        { label: 'GET', value: 'GET' },
        { label: 'POST', value: 'POST' },
        { label: 'PUT', value: 'PUT' },
        { label: 'DELETE', value: 'DELETE' }
      ],
      domainList: [],
      role: { pType: 'g'},
      ruleValidate: {
        v1: [
          { required: true, message: '角色key不能为空', trigger: 'blur' }
        ],
        v5: [
          { required: true, message: '角色名称不能为空', trigger: 'blur' }
        ]
      },

      rom: [], // 角色之前已拥有的菜单id集合
      menuList: [],

      rop: [], // 角色之前已拥有的权限集合
      policyList: [], // 所有的权限集合
    }
  },
  methods: {
    ...mapMutations([
      'closeTag'
    ]),
    ...mapActions([
      'handleDomainList',
      'handleAllMenu',
      'handleRoleOfMenus',
      'handlePolicyAll',
      'handleRoleOfPolicys'
    ]),
    handleSubmit () {
      this.$refs.roleForm.validate(v => {
        if (v) {
          let result = { role: { ...this.role } }
          if (result.role.id) result.role.id = Number(result.role.id)
          // 菜单
          let addMids = []
          let deleteMids = []
          //
          let selectionMenuList = this.$refs.menuTree.getCheckedNodes()
          if (!this.isModify && !this.tRole) { // 新建角色
            selectionMenuList.forEach(sml => addMids.push(sml.id))
            result.addMids = addMids
          } else {
            /** 菜单 */
            if (selectionMenuList.length === 0) { // 修改角色菜单，且将角色菜单清空
              result.deleteMids = this.tRole.menuList.map(e => e.mid)
            } else {
              // 添加的
              selectionMenuList.forEach(e1 => {
                let some = this.tRole.menuList.some(e2 => e1.id === e2.mid)
                if (!some) addMids.push(e1.id)
              })
              // 删除的
              this.tRole.menuList.forEach(e1 => {
                let some = selectionMenuList.some(e2 => e1.mid === e2.id)
                if (!some) deleteMids.push(e1.mid)
              })
              result.addMids = addMids
              result.deleteMids = deleteMids
            }
          }
          this.$emit('submit', result)
        }
      })
    },

    handleReset () {
      this.role = { pType: 'g'}
    },

    handleCloseTag () {
      let c = { name: this.routeName }
      if (this.tRole) c.query = this.tRole
      this.closeTag(c)
    },

    pullDomainList () {
      this.handleDomainList().then(data => {
        if (data) data.forEach(d => this.domainList.push({ value: d, label: d }))
      }).catch(err => console.log('错误：', err))
    },

    pullAllMenu () {
      this.handleAllMenu().then(data => {
        if (data) {
          let rom = []
          if (this.isModify && this.tRole) rom = this.tRole.menuList || []
          data.forEach(d => {
            let children = []
            if (d.children) {
              d.children.forEach(child => {
                let cChecked = false
                if (rom.length > 0) cChecked = rom.some(e => e.mid === child.id)
                children.push({id: child.id, checked: cChecked, title: child.meta.title || ''})
              })
            }
            let pChecked = rom.some(e => e.mid === d.id)
            this.menuList.push({id: d.id, checked: pChecked, title: d.meta.title || '', expand: true, children: children})
          })
        }
      }).catch(err => console.log('错误：', err))
    },

    pullAllPolicy () {
      this.handlePolicyAll().then(data => {
        if (data) {
          data.forEach(d => {
            let checked = false
            if (this.isModify && this.tRole) {
              let rop = this.tRole.policyList
              checked = rop.some(e => e.id === d.id)
            }
            this.policyList.push({id: d.id, checked: checked, title: d.v5 || '', disabled: true})
          })
        }
      }).catch(err => console.log('错误：', err))
    },
  },

  created () {
    if (this.tRole) this.role = { ...this.tRole }
    this.pullDomainList()
    this.pullAllMenu()
    this.pullAllPolicy()
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
