<template>
<section class="some-section-class">
    <img class="some-img-class" :src="$attrs.imgSrc" />
    <h1>see src/.vuepress/components/SampleComponentMarkdownContent.vue</h1>
    <div>
        <slot /> <!-- This is where your markdown content will be inserted -->
    </div>
</section>
</template>

<style scoped>
.some-section-class {
    display: grid;
    border:solid 1px red;
    /* ... more styles ...*/
}
.some-img-class {
    border-radius: 50%;
    /* ... more styles ...*/
}
</style>
