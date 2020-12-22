<template>
  <OnePolicyPage ref="opp" route-name="add_policy_page" @submit="save" />
</template>

<script>
import OnePolicyPage from './one_policy_page'
import { mapActions } from 'vuex'
export default {
  components: {
    OnePolicyPage
  },
  name: 'add_policy_page',
  data () {
    return {}
  },
  methods: {
    ...mapActions([
      'handleCreatePolicy'
    ]),
    save (policy) {
      this.handleCreatePolicy(policy).then(res => {
        this.$refs.opp.handleReset()
        this.$Message.success('创建权限规则成功')
      }).catch(() => this.$Message.error('创建权限规则失败'))
    }
  }
}
</script>
<style scoped>
</style>
