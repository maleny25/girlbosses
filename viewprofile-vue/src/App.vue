<template>
  <div id="app">
    <view-profiles :profiles="profiles" />
  </div>
</template>

<script>
import ViewProfiles from '@/components/ViewProfiles.vue'


const profile = {"name":"a","age":"5"}
  
export default {
  name: 'app',
  components: {
    ViewProfiles
  },
  data() {
    return profile;
  },
  mounted: function() {
    this.displayProfiles;
  },
  methods: {
    displayProfile: displayProfile,
  }
}

async function displayProfile() {
    try {
      const response = await fetch('api/profiles/0')
      const data = await response.json()
      console.log(data)
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
