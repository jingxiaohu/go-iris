<template>
  <OneUserPage ref="oup" route-name="add_user_page" @submit="save" />
</template>

<script>
import OneUserPage from './one_user_page'
import { mapActions } from 'vuex'
export default {
  components: {
    OneUserPage
  },
  name: 'add_user_page',
  data () {
    return {}
  },
  methods: {
    ...mapActions([
      'handleCreateUser'
    ]),
    save (user) {
      this.handleCreateUser(user).then(res => {
        this.$refs.oup.handleReset()
        this.$Message.success('创建用户成功')
      }).catch(() => this.$Message.error('创建用户失败'))
    }
  }
}
</script>
<style scoped>
</style>
