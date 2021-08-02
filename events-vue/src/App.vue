<template>
  <div id="app">
    <events :events="events" />
  </div>
</template>

<script>
import Events from '@/components/Events.vue'

const appData = {
  events: []
}
  
export default {
  name: 'app',
  components: {
    Events
  },
  data() {
    return appData;
  },
  mounted: function() {
    this.getEvents();
  },
  methods: {
    getEvents: getEvents,
  }
}

async function getEvents() {
    try {
      const response = await fetch('api/event')
      const data = await response.json()
      appData.events = data.list
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
  margin-top: 60px;
}
</style>
