<template>
  <div>
    <h1>This is an User page</h1>
    <el-input v-model="id" placeholder="IDを入力して下さい"></el-input>
    <el-input v-model="email" placeholder="メールアドレスを入力して下さい"></el-input>
    <el-input v-model="name" placeholder="名前を入力して下さい"></el-input>
    <el-button @click="GetUser">ユーザーデータの取得</el-button>
    <el-button @click="AddUser">ユーザーデータの登録</el-button>
    <el-button @click="DeleteUser">ユーザーデータの削除</el-button>
    <el-button @click="UpdateUser">ユーザーデータの更新</el-button>
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
      id: '99',
      email: 'test@email',
      name: 'test'
    }
  },
  methods: {
    async GetUser () {
      const response = await axios.get('http://localhost:8080/api/v1/users/1')
      this.user = response.data
    },
    AddUser () {
      axios.post('http://localhost:8080/api/v1/users', {
        id: this.id,
        email: this.email,
        name: this.name
      })
    },
    DeleteUser () {
      axios.delete('http://localhost:8080/api/v1/users/' + this.id)
    },
    UpdateUser () {
      axios.put('http://localhost:8080/api/v1/users/' + this.id, {
        email: this.email,
        name: this.name
      })
    }
  }
}
</script>
