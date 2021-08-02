<template>
  <div id="app">
    <user-form @create:user="createUser" />
    <user-list :users="users" />  

  </div>
</template>

<script>
import UserForm from '@/components/UserForm.vue';
import UserList from '@/components/UserList.vue';
const appData = {
  users: []
}
export default {
  name: 'app',
  components: {
    UserList,
    UserForm
  },
  data() {
    return appData;
  },
  mounted: function() {
    this.getUsers();
  },
  methods: {
    getUsers: getUsers,
    createUser: createUser
  }
}

async function getUsers() {
    try {
      const response = await fetch('api/users')
      const data = await response.json()
      appData.users = data.list
    } 
    catch (error) {
      console.error(error)
    }
}
async function createUser(user) {
  try {
    const response = await fetch('/api/users', {
      method: 'POST',
      body: JSON.stringify(user),
      headers: { "Content-type": "application/json; charset=UTF-8" }
    })
    console.log(response.status)
    await getUsers()
  } 
  catch (error) {
    console.error(error)
  }
}
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  /* margin-top: 60px; */
}
</style>
