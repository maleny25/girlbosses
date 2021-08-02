<template>
  <div id="profile-form" style="text-align: center;">
    <form @submit.prevent="handleSubmit">
    <h2><label>Profile</label></h2>

    <p>&nbsp;</p>
    <table style="height: 232px; width: 98.4134%; border-collapse: collapse; border-style: hidden;" border="0">
      <tbody>
      <tr style="height: 18px;">
      <td style="width: 43.2991%; height: 18px; text-align: right;">Name* :</td>
      <td style="width: 93.5772%; height: 18px; text-align: left;"><input v-model="profile.name" type="text" /></td>
      </tr>
      <tr style="height: 18px;">
      <td style="width: 43.2991%; height: 18px; text-align: right;">Age :</td>
      <td style="width: 93.5772%; height: 18px; text-align: left;"><input v-model="profile.age" type="text" /></td>
      </tr>
      <tr style="height: 18px;">
      <td style="width: 43.2991%; height: 18px; text-align: right;">Gender :</td>
      <td style="width: 93.5772%; height: 18px; text-align: left;"><input v-model="profile.gender" type="text" /></td>
      </tr>
      <tr style="height: 18px;">
      <td style="width: 43.2991%; height: 18px; text-align: right;">Race :</td>
      <td style="width: 93.5772%; height: 18px; text-align: left;"><input v-model="profile.race" type="text" /></td>
      </tr>
      <tr style="height: 18px;">
      <td style="width: 43.2991%; height: 18px; text-align: right;">&nbsp;</td>
      <td style="width: 93.5772%; height: 18px; text-align: left;">&nbsp;</td>
      </tr>
      <tr style="height: 18px;">
      <td style="width: 43.2991%; height: 18px; text-align: right;">University* :</td>
      <td style="width: 93.5772%; height: 18px; text-align: left;"><input v-model="profile.university" type="text" /></td>
      </tr>
      <tr style="height: 18px;">
      <td style="width: 43.2991%; height: 18px; text-align: right;">Major* :</td>
      <td style="width: 93.5772%; height: 18px; text-align: left;"><input v-model="profile.major" type="text" /></td>
      </tr>
      <tr style="height: 18px;">
      <td style="width: 43.2991%; height: 18px; text-align: right;">Minor :</td>
      <td style="width: 93.5772%; height: 18px; text-align: left;"><input v-model="profile.minor" type="text" /></td>
      </tr>
      <tr style="height: 21px;">
      <td style="width: 43.2991%; height: 21px; text-align: right;">Grad Year* :</td>
      <td style="width: 93.5772%; height: 21px; text-align: left;"><input v-model="profile.gradyear" type="text" /></td>
      </tr>
      <tr style="height: 18px;">
      <td style="width: 43.2991%; height: 18px; text-align: right;">&nbsp;</td>
      <td style="width: 93.5772%; height: 18px; text-align: left;">&nbsp;</td>
      </tr>
      <tr style="height: 49px;">
      <td style="width: 43.2991%; height: 49px; text-align: right;">
      <p>Description* :</p>
      </td>
      <td style="width: 93.5772%; height: 49px; text-align: left;">
      <p><input v-model="profile.description" type="text" /></p>
      </td>
      </tr>
      </tbody>
    </table>
      <button>Submit</button>
    </form>
  </div>
</template>

<script>
export default {
  name: 'profile-form',
  data() {
    return {
      error: false,
      submitting: false,
      success: false,
      profile: {
        name: '',
        university: '',
        major: '',
      }
    }
  },
  computed: {
    invalidName() {
      return this.profile.name === ''
    },
    invalidUniversity() {
      return this.profile.university === ''
    },
    // invalidMajor() {
    //   return this.profile.major === ''
    // },
    // invalidGradYear() {
    //   return this.profile.gradyear === ''
    // },
    // invalidDescription() {
    //   return this.profile.description === ''
    // },
  },
  methods: {
    handleSubmit() {
      this.clearStatus()
      this.submitting = true
      if (this.invalidName || this.invalidUniversity ){//|| this.invalidMajor || this.invalidGradYear || this.invalidDescription) {
        this.error = true
        return
      }
      this.$emit('add:profile', this.profile)
      this.$refs.first.focus()
      this.profile = {
        name: '',
        university: '',
        major: '',
      }
      this.success = true
      this.error = false
      this.submitting = false
    },
    clearStatus() {
      this.success = false
      this.error = false
    }
  }}
</script>

<style scoped>
form {
  margin-bottom: 2rem;
}
[class*="-message"] {
  font-weight: 500;
}
.error-message {
  color: #d33c40;
}
.success-message {
  color: #32a95d;
}
</style>
