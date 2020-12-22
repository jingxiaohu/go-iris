<template>
  <OneMenuPage ref="omp" route-name="modify_menu_page" :t-menu="menu" @submit="save"/>
</template>

<script>
import OneMenuPage from './one_menu_page'
import { mapActions } from 'vuex'
export default {
  components: {
    OneMenuPage
  },
  name: 'modify_menu_page',
  data () {
    return {
      menu: null
    }
  },
  methods: {
    ...mapActions([
      'handleUpdateMenu'
    ]),
    save (menu) {
      // console.log(menu)
      // return
      this.handleUpdateMenu(menu).then(res => {
        this.$refs.omp.handleCloseTag()
        this.$Message.success('更新菜单信息成功')
      }).catch(() => this.$Message.error('更新菜单信息失败'))
    }
  },
  created () {
    this.menu = JSON.parse(sessionStorage.getItem('selectionMenu') || '{}')
  }
}
</script>
<style scoped>
</style>
