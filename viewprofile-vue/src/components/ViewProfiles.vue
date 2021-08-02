<template>
<div id="view-profiles" style="text-align: center;">
    <h2><label>Your Profile</label></h2>
    
<table style="width: 100%; border-collapse: collapse; border-style: hidden;" border="0">
<tbody>
<tr style="text-align: center;">
<td>
<p style="text-align: right;">Name:</p>
<p style="text-align: right;">Age:</p>
<p style="text-align: right;">Gender:</p>
<p style="text-align: right;">Race:</p>
<p style="text-align: right;">&nbsp;</p>
<p style="text-align: right;">University:</p>
<p style="text-align: right;">Major:</p>
<p style="text-align: right;">Minor:</p>
<p style="text-align: right;">Grad Year:</p>
<p style="text-align: right;">&nbsp;</p>
<p style="text-align: right;">Description:</p>
</td>
<td style="padding-left: 40px;">
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
<p style="text-align: left;">Name:</p>
<p style="text-align: left;">Age:</p>
<p style="text-align: left;">Gender:</p>
<p style="text-align: left;">Race:</p>
<p style="text-align: left;">&nbsp;</p>
<p style="text-align: left;">University:</p>
<p style="text-align: left;">Major:</p>
<p style="text-align: left;">Minor:</p>
<p style="text-align: left;">Grad Year:</p>
<p style="text-align: left;">&nbsp;</p>
<p style="text-align: left;">Description:</p>
</td>
</tr>
</tbody>
</table>
</div>

</template>

<script>


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
