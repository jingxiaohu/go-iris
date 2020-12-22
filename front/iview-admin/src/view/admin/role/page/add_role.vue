<template>
  <OneRolePage ref="orp" route-name="add_role_page" @submit="save" />
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
    return {}
  },
  methods: {
    ...mapActions([
      'handleCreateRole'
    ]),
    save (role) {
      this.handleCreateRole(role).then(res => {
        this.$refs.orp.handleReset()
        this.$Message.success('创建角色成功')
      }).catch((err) => {
        console.log(err)
        this.$Message.error('创建角色失败')
      })
    }
  },
  created () {
  }
}
</script>
<style scoped>
</style>
