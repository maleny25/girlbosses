<template>
  <div id="app">
    <profile-form @add:profile="addProfile" />
    <profile-list :profiles="profiles" />
  </div>
</template>

<script>
import ProfileForm from '@/components/ProfileForm.vue'
import ProfileList from '@/components/ProfileList.vue'


const appData = {
  profiles: []
}
  
export default {
  name: 'app',
  components: {
    ProfileForm,
    ProfileList
  },
  data() {
    return appData;
  },
  mounted: function() {
    this.getProfiles();
  },
  methods: {
    getProfiles: getProfiles,
    addProfiles: addProfiles
  }
}

async function getProfiles() {
    try {
      const response = await fetch('api/profiles')
      const data = await response.json()
      appData.profiles = data.list
    } 
    catch (error) {
      console.error(error)
    }
}

async function addProfiles(profile) {
  try {
    const response = await fetch('/api/profiles', {
      method: 'POST',
      body: JSON.stringify(profile),
      headers: { "Content-type": "application/json; charset=UTF-8" }
    })
    console.log(response.status)
    await getProfiles()
  } 
  catch (error) {
    console.error(error)
  }
}

</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  margin-top: 50px;
}
</style>
