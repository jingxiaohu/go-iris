<template>
  <div>
    <Form ref="menuForm" :model="menu" :rules="ruleValidate" label-position="right" :label-width="130">
      <Row :gutter="2">
        <h4>菜单信息</h4>
        <Col span="12" push="3">
          <FormItem prop="path" label="路径：">
            <Input type="text" v-model="menu.path" placeholder="请输入路径名称" clearable style="width: auto"/>
          </FormItem>
          <FormItem prop="name" label="路径名称：">
            <Input type="text" v-model="menu.name" placeholder="请输入路径名称" clearable style="width: auto"/>
          </FormItem>
          <FormItem label="模块：">
            <Input type="text" v-model="menu.modular" :disabled="!menu.isSub" placeholder="请输入模块名称" clearable style="width: auto"/>
          </FormItem>
          <FormItem label="组件：">
            <Input type="text" v-model="menu.component" :disabled="!menu.isSub" placeholder="请输入组件名称" clearable style="width: auto"/>
          </FormItem>
        </Col>
        <Col span="12">
          <FormItem label="子菜单：">
            <i-switch size="large" v-model="menu.isSub" @on-change="isSubChange">
              <span slot="open">是</span>
              <span slot="close">否</span>
            </i-switch>
          </FormItem>
          <FormItem label="父级菜单：">
            <Select v-model="menu.parentId" :disabled="!menu.isSub" @on-open-change="loadParentMenus"  style="width: 188px">
              <Option v-for="item in parentMenuList" :value="item.value" :key="item.value">{{ item.label }}</Option>
            </Select>
            {{menu.parentId}}
          </FormItem>
          <FormItem label="菜单顺序：">
            <Input type="number" v-model="menu.sort" number placeholder="请输入菜单顺序" style="width: auto"/>
          </FormItem>
        </Col>
      </Row>

      <h4>菜单元信息(meta)</h4>
      <Row :gutter="2">
        <Col span="12" push="3">
          <FormItem prop="meta.title" label="菜单名称：">
            <Input type="text" v-model="menu.meta.title" placeholder="请输入菜单名称" clearable style="width: auto"/>
          </FormItem>
          <FormItem @click.native="icon.modal=!icon.modal" prop="meta.icon" label="菜单图标：">
            <Input type="text" v-model="menu.meta.icon" placeholder="请输入菜单图标" readonly :prefix="menu.meta.icon" style="width: auto"/>
          </FormItem>
          <FormItem label="钩子函数：">
            <Input type="text" v-model="menu.meta.beforeCloseName" placeholder="请输入钩子函数" clearable style="width: auto"/>
          </FormItem>
          <FormItem label="外部连接：">
            <Input type="text" v-model="menu.meta.href" placeholder="请输入钩子函数" clearable style="width: auto"/>
          </FormItem>
          <FormItem label="权限数组：">
            <Input disabled type="text" v-model="menu.meta.access" placeholder="请输入权限数组" clearable style="width: auto"/>
          </FormItem>
        </Col>
        <Col span="12">
          <FormItem label="父级菜单显示：">
            <i-switch size="large" v-model="menu.meta.showAlways">
              <span slot="open">ON</span>
              <span slot="close">OFF</span>
            </i-switch>
          </FormItem>
          <FormItem label="关闭面包屑导航：">
            <i-switch size="large" v-model="menu.meta.hideInBread">
              <span slot="open">ON</span>
              <span slot="close">OFF</span>
            </i-switch>
          </FormItem>
          <FormItem label="不显示菜单：">
            <i-switch size="large" v-model="menu.meta.hideInMenu">
              <span slot="open">ON</span>
              <span slot="close">OFF</span>
            </i-switch>
          </FormItem>
          <FormItem label="不缓存页面：">
            <i-switch size="large" v-model="menu.meta.notCache">
              <span slot="open">ON</span>
              <span slot="close">OFF</span>
            </i-switch>
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

    <Modal
      draggable
      v-model="icon.modal"
      width="1200"
      title='图标集合'>
      <Row :gutter="0" v-for="(row, i) in icon.icons" :key="i">
        <Col span="4" v-for="(icon, j) in row" :key="`${i}${j}`">
          <Button @click="selectIcon(icon.icon)" type="text"><Icon size="20" :type="icon.icon" /> {{icon.icon}}</Button>
        </Col>
        <br />
      </Row>
      <span slot="footer"></span>
    </Modal>
  </div>
</template>

<script>
import { mapMutations, mapActions } from 'vuex'
import {icons} from '@/assets/icons'
export default {
  props: {
    tMenu: {
      type: Object,
      default: () => null
    },
    routeName: String
  },
  name: 'OneMenuPage',
  data () {
    return {
      menu: { meta: {} },
      ruleValidate: {
        path: [
          { required: true, message: '路径不能为空', trigger: 'blur' }
        ],
        name: [
          { required: true, message: '路径名称不能为空', trigger: 'blur' }
        ],
        // modular: [
        //   { required: true, message: '模块不能为空', trigger: 'blur' }
        // ],
        // component: [
        //   { required: true, message: '组件不能为空', trigger: 'blur' }
        // ],
        // meta
        'meta.title': [
          { required: true, message: '菜单名称不能为空', trigger: 'blur' }
        ],
        'meta.icon': [
          { required: true, message: '图标不能为空', trigger: 'blur' }
        ]
      },
      icon: {
        icons: icons,
        modal: false
      },
      parentMenuList: []
    }
  },
  methods: {
    ...mapMutations([
      'closeTag'
    ]),
    ...mapActions([
      'handleParentMenus'
    ]),
    handleSubmit () {
      this.$refs.menuForm.validate(v => {
        if (v) {
          if (this.menu.isSub) {
            if (!this.menu.modular) {
              this.$Message.error('模块不能为空')
              return
            }
            if (!this.menu.component) {
              this.$Message.error('组件不能为空')
              return
            }
            if (!this.menu.parentId) {
              this.$Message.error('父级菜单不能为空')
              return
            }
          }
          this.$emit('submit', this.menu)
        }
      })
    },
    handleReset () {
      if (this.tMenu) this.menu = { ...this.tMenu }
      else this.menu = { meta: {} }
      this.parentMenuList = []
    },

    handleCloseTag () {
      this.closeTag({ name: this.routeName })
      sessionStorage.removeItem('selectionMenu')
    },

    selectIcon (icon) {
      this.icon.modal = false
      this.$set(this.menu.meta, 'icon', icon)
    },

    isSubChange (ch) {
      if (!ch) {
        delete this.menu.modular
        delete this.menu.component
      }
    },

    loadParentMenus (ch) {
      if (ch && this.parentMenuList.length < 1) {
        this.handleParentMenus().then(data => {
          if (data) {
            if (this.menu.isSub) this.parentMenuList = []
            data.forEach(e => {
              let title = '未知'
              if (e.meta) title = e.meta.title
              this.parentMenuList.push({ value: e.id, label: title })
            })
          }
        }).catch(err => console.log('获取所有父菜单数据错误：', err))
      }
    }
  },
  created () {
    // this.pullParentMenus()
    if (this.tMenu) this.menu = { ...this.tMenu }
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
