<template>
  <d2-container>
    <d2-crud
      :columns="columns"
      :data="data"
      :rowHandle="rowHandle"
      @custom-emit-1="handleCustomEvent"
    >
    </d2-crud>
  </d2-container>
</template>

<script>
import moment from 'moment'
import D2Crud from '@d2-projects/d2-crud'
import { QueryUserByLimitOffset } from '@api/user'
export default {
  components: { D2Crud },
  data() {
    return {
      pagination: {
        currentPage: 1,
        pageSize: 15,
        total: 0,
      },
      columns: [
        {
          title: 'ID',
          key: 'id',
          width: '50px',
        },
        {
          title: '用户名',
          key: 'username',
        },
        {
          title: '电话号码',
          key: 'tel',
        },
        {
          title: 'ip',
          key: 'ip',
        },
        {
          title: '创建时间',
          key: 'created_at',
          formatter: function(row) {
            if (row.created_at == undefined) {
              return ''
            }
            return moment(row.created_at).format('YYYY-MM-DD HH:mm:ss')
          },
        },
      ],
      data: [],
      rowHandle: {
        custom: [
          {
            icon: 'el-icon-view',
            type: 'info',
            text: '详情',
            size: 'small',
            emit: 'custom-emit-1',
          },
        ],
      },
    }
  },
  mounted() {
    this.fetchData()
  },
  methods: {
    // 分页
    paginationCurrentChange(currentPage) {
      this.pagination.currentPage = currentPage
      this.fetchData()
    },
    // 获取数据
    fetchData() {
      QueryUserByLimitOffset(
        this.pagination.pageSize,
        this.pagination.currentPage
      )
        .then(async (res) => {
          this.data = res.users
          this.pagination.total = res.count
        })
        .catch((err) => {
          Promise.reject(err)
        })
    },
    handleCustomEvent({ index, row }) {
      console.log(index)
      console.log(row)
    },
  },
}
</script>
