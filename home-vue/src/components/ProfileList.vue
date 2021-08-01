<template>
  <div id="profile-list">
    <h1>Profile List</h1>
    <table>
      <thead>
        <tr>
          <th>Name</th>
          <th>University</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="profile in profiles" :key="profile.id">
          <td>{{ profile.name }}</td>
          <td>{{ profile.university }}</td>
        </tr>
      </tbody>
    </table>
  </div>
</template>

<script>
export default {
  name: 'profile-list',
  props: {
    profiles: Array,
  },
  data() {
    return null;
  }
}
</script>

<style scoped>
</style>