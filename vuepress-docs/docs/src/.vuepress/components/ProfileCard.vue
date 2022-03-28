<template>
<section class="profile-card">
    <img class="profile-img" :src="$attrs.imgSrc" />
    <h1>Sample component with rich content</h1>
    <div>
        <slot /> <!-- This is where your markdown content will be inserted -->
    </div>
</section>
</template>

<style scoped>
.profile-card {
    display: grid;
    /* ... more styles ...*/
}

.profile-img {
    border-radius: 50%;
    /* ... more styles ...*/
}
</style>
