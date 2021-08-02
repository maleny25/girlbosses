<template>
  <nav>
    <ul>
      <li><a href="/">Home</a></li>
      <li><a href="/">Foo</a></li>
    </ul>
  </nav>
</template>

<style scoped>
  ul {
    list-style-type: none;
    margin: 0;
    padding: 0;
    display: flex;
    justify-content: center;
  }
  
  li:not(:last-child) {
    margin-right: 16px;
  }
</style>