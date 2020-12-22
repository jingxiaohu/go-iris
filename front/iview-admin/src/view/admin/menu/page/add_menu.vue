<template>
  <OneMenuPage ref="orp" route-name="add_menu_page" @submit="save" />
</template>

<script>
import OneMenuPage from './one_menu_page'
import { mapActions } from 'vuex'
export default {
  components: {
    OneMenuPage
  },
  name: 'add_menu_page',
  data () {
    return {}
  },
  methods: {
    ...mapActions([
      'handleCreateMenu'
    ]),
    save (role) {
      this.handleCreateMenu(role).then(res => {
        this.$refs.orp.handleReset()
        this.$Message.success('创建菜单成功')
      }).catch(() => this.$Message.error('创建菜单失败'))
    }
  },
  created () {}
}
</script>
<style scoped>
</style>
