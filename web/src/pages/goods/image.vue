<template>
  <div>
    <el-image
      style="width: 50px; height:50px"
      v-for="url in img"
      :key="url"
      :src="url"
      :preview-src-list="imgs"
    >
    </el-image>
  </div>
</template>

<script>
export default {
  data () {
    return {
      imgs: [],
      img: [],
      data: []
    }
  },
  props: {
    value: ''
  },
  mounted () {
    this.fetchData()
  },
  methods: {
    fetchData () {
      this.data = this.value
      if (this.data.length > 0) {
        this.img.push(this.data[0].goods_image_files)
      } else {
        this.img.push('')
      }
      this.data.forEach(itme => {
        this.imgs.push(itme.goods_image_files)
      });
    },
  }
}
</script>