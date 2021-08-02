<template>
  <div id="user-list">
    <h1>User List</h1>
    <table>
      <thead>
        <tr>
          <th>Username</th>
          <th>Password</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="user in users" :key="user.id">
          <td>{{ user.username }}</td>
          <td>{{ user.password }}</td>
        </tr>
      </tbody>
    </table>
  </div>
</template>

<script>
export default {
  name: 'user-list',
  props: {
    users: Array,
  },
  data() {
    return null;
  }
}
</script>

<style scoped>
</style>