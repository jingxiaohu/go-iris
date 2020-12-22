<template>
  <OneUserPage ref="oup" route-name="modify_user_page" :t-user="user" @submit="save"/>
</template>

<script>
import OneUserPage from './one_user_page'
import { mapActions } from 'vuex'
export default {
  components: {
    OneUserPage
  },
  name: 'modify_user_page',
  data () {
    return {
      user: null
    }
  },
  methods: {
    ...mapActions([
      'handleUpdateUser'
    ]),
    save (user) {
      this.handleUpdateUser(user).then(res => {
        this.$refs.oup.handleCloseTag()
        this.$Message.success('更新用户信息成功')
      }).catch(() => this.$Message.error('更新用户信息失败'))
    }
  },
  created () {
    this.user = JSON.parse(sessionStorage.getItem('selectionUser') || '{}')
  }
}
</script>
<style scoped>
</style>
