<template>
  <div id="profile-list">
    <h1>Profile List</h1>
    <table>
      <thead>
        <tr>
          <th>Name</th>
          <th>Age</th>
          <th>Gender</th>
          <th>Race</th>
          <th>University</th>
          <th>Major</th>
          <th>Minor</th>
          <th>Grad_Year</th>
          <th>Description</th>

        </tr>
      </thead>
      <tbody>
        <tr v-for="profile in profiles" :key="profile.id">
          <td>{{ profile.name }}</td>
          <td>{{ profile.age }}</td>
          <td>{{ profile.gender }}</td>
          <td>{{ profile.race }}</td>
          <td>{{ profile.university }}</td>
          <td>{{ profile.major }}</td>
          <td>{{ profile.minor }}</td>
          <td>{{ profile.gradyear }}</td>
          <td>{{ profile.description }}</td>

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