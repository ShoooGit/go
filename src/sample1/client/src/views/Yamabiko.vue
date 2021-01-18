<template>
  <div>
    <h1>This is an Yamabiko page</h1>
    <!-- <input v-model="message" placeholder="Say Yahho"> -->
    <button @click="SampleAPI">ユーザーデータの取得</button>
    <!-- <p>Yamabiko : {{ yamabiko }}</p> -->
  </div>
</template>

<script>
const axios = require('axios').create()
export default {
  name: 'Yamabiko',
  data () {
    return {
      user: []
    }
  },
  methods: {
    async SampleAPI () {
      const response = await axios.get('http://localhost:8080/api/v1/users/1')
      console.log(response.data)
      this.user = response.data
      console.log(response.data)
    }
  }
}
</script>
