<template>
  <div>
    <h1>This is an User page</h1>
    <!-- <input v-model="message" placeholder="Say Yahho"> -->
    <!-- <button @click="GetUser">ユーザーデータの取得</button> -->
    <button @click="AddUser">ユーザーデータの登録</button>
    <p>ユーザー : {{ user }}</p>
  </div>
</template>

<script>
const axios = require('axios').create()
export default {
  name: 'User',
  data () {
    return {
      user: []
    }
  },
  methods: {
    // async GetUser () {
    //   const response = await axios.get('http://localhost:8080/api/v1/users/1')
    //   this.user = response.data
    // },
    AddUser () {
      axios.post('http://localhost:8080/api/v1/users', {
        email: 'Non error debitis ut quia.',
        id: 'XRQ85mtXnINISH25zfM0m5RlC6L2',
        name: 'Sequi natus repellat illum eum.'
      })
    }
  }
}
</script>
