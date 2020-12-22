<template>
  <OnePolicyPage ref="opp" route-name="modify_policy_page" :t-policy="policy" @submit="save"/>
</template>

<script>
import OnePolicyPage from './one_policy_page'
import { mapActions } from 'vuex'
export default {
  components: {
    OnePolicyPage
  },
  name: 'modify_policy_page',
  data () {
    return {
      policy: null
    }
  },
  methods: {
    ...mapActions([
      'handleUpdatePolicy'
    ]),
    save (policy) {
      this.handleUpdatePolicy(policy).then(res => {
        this.$refs.opp.handleCloseTag()
        this.$Message.success('更新策略规则成功')
      }).catch(() => this.$Message.error('更新策略规则失败'))
    }
  },
  created () {
    this.policy = this.$route.query
    console.log(this.policy)
  }
}
</script>
<style scoped>
</style>
