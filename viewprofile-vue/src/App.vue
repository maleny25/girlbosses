<template>
  <div id="app">
    <view-profiles :profiles="profiles" />
  </div>
</template>

<script>
import ViewProfiles from '@/components/ViewProfiles.vue'


const appData = {
  profiles: []
}
  
export default {
  name: 'app',
  components: {
    ViewProfiles
  },
  data() {
    return appData;
  },
  mounted: function() {
    this.getProfiles();
  },
  methods: {
    getProfiles: getProfiles,
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


</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  margin-top: 50px;
}
</style>
