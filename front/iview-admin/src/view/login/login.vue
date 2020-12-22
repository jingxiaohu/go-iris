<template>
  <div class="login">
    <div class="login-con">
      <Card icon="log-in" title="欢迎登录go iris" :bordered="false">
        <div class="form-con">
          <login-form @on-success-valid="handleSubmit"></login-form>
          <p class="login-tip">输入任意用户名和密码即可</p>
        </div>
      </Card>
    </div>
  </div>
</template>

<script>
import LoginForm from '_c/login-form'
import { mapActions } from 'vuex'
export default {
  components: {
    LoginForm
  },
  methods: {
    ...mapActions([
      'handleLogin',
      'getUserInfo'
    ]),
    handleSubmit (user) {
      localStorage.clear()
      this.handleLogin(user).then(res => {
        this.getUserInfo(user.id).then(res => {
          if (this.$router.options.routes[2]) {
            console.log(this.$router.options.routes[2])
            this.$router.options.routes[2].meta.hideInMenu = true
            // this.$router.options.routes[2].meta.hideInBread = true
          }
          console.log('----', this.$router.options.routes)
          this.$router.push({
            name: this.$config.homeName
          })
        })
      })
    }
  }
}
</script>

<style lang="less">
  @import './login.less';
</style>
