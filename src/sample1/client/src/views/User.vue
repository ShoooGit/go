<template>
  <div>
    <h1>This is an User page</h1>
    <div class="form">
      <el-input v-model="id" placeholder="IDを入力して下さい"></el-input>
      <el-input v-model="email" placeholder="メールアドレスを入力して下さい"></el-input>
      <el-input v-model="name" placeholder="名前を入力して下さい"></el-input>
    </div>
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
      const response = await axios.get('http://localhost:8080/api/v1/users/' + this.id)
      this.user = response.data
    },
    async AddUser () {
      const response = await axios.post('http://localhost:8080/api/v1/users', {
        id: this.id,
        email: this.email,
        name: this.name
      })
      alert(response.data)
    },
    DeleteUser () {
      axios.delete('http://localhost:8080/api/v1/users/' + this.id)
    },
    async UpdateUser () {
      const response = await axios.put('http://localhost:8080/api/v1/users/' + this.id, {
        email: this.email,
        name: this.name
      })
      this.user = response.data
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
