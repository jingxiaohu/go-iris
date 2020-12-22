<template>
  <OneRolePage ref="orp" route-name="modify_role_page" :t-role="routeQuery" :is-modify="true" @submit="save"/>
</template>

<script>
import OneRolePage from './one_role_page'
import { mapActions } from 'vuex'
export default {
  components: {
    OneRolePage
  },
  name: 'add_role_page',
  data () {
    return {
      routeQuery: null
    }
  },
  methods: {
    ...mapActions([
      'handleUpdateRole'
    ]),
    save (role) {
      this.handleUpdateRole(role).then(res => {
        this.$refs.orp.handleCloseTag()
        this.$Message.success('更新角色信息成功')
      }).catch(() => this.$Message.error('更新角色信息失败'))
    }
  },
  created () {
    this.routeQuery = this.$route.query
  }
}
</script>
<style scoped>
</style>
