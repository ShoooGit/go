<template>
  <div>
    <h1>This is an Ha page</h1>
    <div class="form">
      <el-button @click="DrawCard">カードを引く</el-button>
      <p>カード : {{ card }}</p>
    </div>
  </div>
</template>

<script>
const axios = require('axios').create()
export default {
  name: 'Ha',
  data () {
    return {
      card: ''
    }
  },
  methods: {
    async DrawCard () {
      const response = await axios.get('http://localhost:8080/api/v1/ha/draw')
      this.card = response.data
    }
  }
}
</script>
<style scoped>
.form {
    width: 30%;
    margin: auto;
}
</style>
