<template>
  <div>
    <h1>This is an User page</h1>
    <el-input v-model="email" placeholder="メールアドレスを入力して下さい"></el-input>
    <el-input v-model="name" placeholder="名前を入力して下さい"></el-input>
    <el-button @click="GetUser">ユーザーデータの取得</el-button>
    <el-button @click="AddUser">ユーザーデータの登録</el-button>
    <p>ユーザー : {{ user }}</p>
  </div>
</template>

<script>
const axios = require('axios').create()
export default {
  name: 'User',
  data () {
    return {
      user: [],
      email: '',
      name: ''
    }
  },
  methods: {
    async GetUser () {
      const response = await axios.get('http://localhost:8080/api/v1/users/1')
      this.user = response.data
    },
    AddUser () {
      axios.post('http://localhost:8080/api/v1/users', {
        id: '99',
        email: 'test@email',
        name: 'test'
      })
    }
  }
}
</script>
