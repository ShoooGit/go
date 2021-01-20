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
        email: 'test@email',
        id: '99',
        name: 'test'
      })
    }
  }
}
</script>
