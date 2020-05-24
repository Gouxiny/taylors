<template>
  <div>
    <div class="search-term">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
        <el-form-item label="市值(最小)">
          <el-input placeholder="路径" v-model="searchInfo.marketCapitalMin"></el-input>
        </el-form-item>
        <el-form-item label="市值(最大)">
          <el-input placeholder="路径" v-model="searchInfo.marketCapitalMax"></el-input>
        </el-form-item>
        <el-form-item label="涨幅(最小)">
          <el-input placeholder="描述" v-model="searchInfo.percentMin"></el-input>
        </el-form-item>
        <el-form-item label="涨幅(最大)">
          <el-input placeholder="描述" v-model="searchInfo.percentMax"></el-input>
        </el-form-item>
<!--        <el-form-item label="量比(最小)">-->
<!--          <el-input placeholder="api组" v-model="searchInfo.apiGroup"></el-input>-->
<!--        </el-form-item>-->
<!--        <el-form-item label="量比(最大)">-->
<!--          <el-input placeholder="api组" v-model="searchInfo.apiGroup"></el-input>-->
<!--        </el-form-item>-->
        <el-form-item>
          <el-button @click="onSubmit" type="primary">查询</el-button>
        </el-form-item>
<!--        <el-form-item>-->
<!--          <el-button @click="openDialog('addFilterTopApi')" type="primary">保存条件</el-button>-->
<!--        </el-form-item>-->
      </el-form>
    </div>
    <el-table :data="tableData" @sort-change="sortChange" border stripe>
      <el-table-column label="名称" min-width="60" prop="symbol" sortable="custom"></el-table-column>
      <el-table-column label="市值" min-width="150" prop="market_capital" sortable="custom"></el-table-column>
      <el-table-column label="涨幅" min-width="150" prop="percent" sortable="custom"></el-table-column>
      <el-table-column label="涨停" min-width="150" prop="high" sortable="custom"></el-table-column>
      <el-table-column label="跌停" min-width="150" prop="limit_down" sortable="custom"></el-table-column>
      <el-table-column label="最高" min-width="150" prop="chg" sortable="custom"></el-table-column>
      <el-table-column label="最低" min-width="150" prop="low" sortable="custom"></el-table-column>
      <el-table-column label="今开" min-width="150" prop="open" sortable="custom"></el-table-column>
      <el-table-column label="昨收" min-width="150" prop="last_close" sortable="custom"></el-table-column>
      <el-table-column label="成交量" min-width="150" prop="volume" sortable="custom"></el-table-column>
      <el-table-column label="成交额" min-width="150" prop="amount" sortable="custom"></el-table-column>
<!--        <template slot-scope="scope">-->
<!--          <div>-->
<!--            {{scope.row.method}}-->
<!--            <el-tag-->
<!--                    :key="scope.row.methodFiletr"-->
<!--                    :type="scope.row.method|tagTypeFiletr"-->
<!--                    effect="dark"-->
<!--                    size="mini"-->
<!--            >{{scope.row.method|methodFiletr}}</el-tag>-->
<!--            &lt;!&ndash; {{scope.row.method|methodFiletr}} &ndash;&gt;-->
<!--          </div>-->
<!--        </template>-->
<!--      </el-table-column>-->

      <el-table-column fixed="right" label="操作" width="200">
        <template slot-scope="scope">
          <el-button @click="editApi(scope.row)" size="small" type="text">编辑</el-button>
          <el-button @click="deleteApi(scope.row)" size="small" type="text">删除</el-button>
        </template>
      </el-table-column>
    </el-table>
  </div>
</template>


<script>
  // 获取列表内容封装在mixins内部  getTableData方法 初始化已封装完成 条件搜索时候 请把条件安好后台定制的结构体字段 放到 this.searchInfo 中即可实现条件搜索

  import {
    getApiById,
    getTopList,
    createApi,
    updateApi,
    deleteApi
  } from '@/api/stock'
  import infoList from '@/components/mixins/infoList'
  import { toSQLLine } from '@/utils/stringFun'
  const methodOptions = [
    {
      value: 'POST',
      label: '创建',
      type: 'success'
    },
    {
      value: 'GET',
      label: '查看',
      type: ''
    },
    {
      value: 'PUT',
      label: '更新',
      type: 'warning'
    },
    {
      value: 'DELETE',
      label: '删除',
      type: 'danger'
    }
  ]

  export default {
    name: 'Api',
    mixins: [infoList],
    data() {
      return {
        listApi: getTopList,
        dialogFormVisible: false,
        dialogTitle: '新增Api',
        form: {
          path: '',
          apiGroup: '',
          method: '',
          description: ''
        },
        methodOptions: methodOptions,
        type: '',
        rules: {
          path: [{ required: true, message: '请输入api路径', trigger: 'blur' }],
          apiGroup: [
            { required: true, message: '请输入组名称', trigger: 'blur' }
          ],
          method: [
            { required: true, message: '请选择请求方式', trigger: 'blur' }
          ],
          description: [
            { required: true, message: '请输入api介绍', trigger: 'blur' }
          ]
        }
      }
    },
    methods: {
      // 排序
      sortChange({ prop, order }) {
        if (prop) {
          this.searchInfo.orderKey = toSQLLine(prop)
          this.searchInfo.desc = order == 'descending'
        }
        this.getTableData()
      },
      //条件搜索前端看此方法
      onSubmit() {
        this.page = 1
        this.pageSize = 10
        this.getTableData()
      },
      initForm() {
        this.$refs.apiForm.resetFields()
        this.form= {
          path: '',
          apiGroup: '',
          method: '',
          description: ''
        }
      },
      closeDialog() {
        this.initForm()
        this.dialogFormVisible = false
      },
      openDialog(type) {
        switch (type) {
          case 'addApi':
            this.dialogTitlethis = '新增Api'
            break
          case 'edit':
            this.dialogTitlethis = '编辑Api'
            break
          default:
            break
        }
        this.type = type
        this.dialogFormVisible = true
      },
      async editApi(row) {
        const res = await getApiById({ id: row.ID })
        this.form = res.data.api
        this.openDialog('edit')
      },
      async deleteApi(row) {
        this.$confirm('此操作将永久删除所有角色下该菜单, 是否继续?', '提示', {
          confirmButtonText: '确定',
          cancelButtonText: '取消',
          type: 'warning'
        })
                .then(async () => {
                  const res = await deleteApi(row)
                  if (res.code == 0) {
                    this.$message({
                      type: 'success',
                      message: '删除成功!'
                    })
                    this.getTableData()
                  }
                })
                .catch(() => {
                  this.$message({
                    type: 'info',
                    message: '已取消删除'
                  })
                })
      },
      async enterDialog() {
        this.$refs.apiForm.validate(async valid => {
          if (valid) {
            switch (this.type) {
              case 'addApi':
              {
                const res = await createApi(this.form)
                if (res.code == 0) {
                  this.$message({
                    type: 'success',
                    message: '添加成功',
                    showClose: true
                  })
                }
                this.getTableData()
                this.closeDialog()
              }

                break
              case 'edit':
              {
                const res = await updateApi(this.form)
                if (res.code == 0) {
                  this.$message({
                    type: 'success',
                    message: '编辑成功',
                    showClose: true
                  })
                }
                this.getTableData()
                this.closeDialog()
              }
                break
              default:
              {
                this.$message({
                  type: 'error',
                  message: '未知操作',
                  showClose: true
                })
              }
                break
            }
          }
        })
      }
    },
    filters: {
      methodFiletr(value) {
        const target = methodOptions.filter(item => item.value === value)[0]
        // return target && `${target.label}(${target.value})`
        return target && `${target.label}`
      },
      tagTypeFiletr(value) {
        const target = methodOptions.filter(item => item.value === value)[0]
        return target && `${target.type}`
      }
    },
    created(){
      this.getTableData()
    }
  }
</script>
<style scoped lang="scss">
  .button-box {
    padding: 10px 20px;
  .el-button {
    float: right;
  }
  }
  .el-tag--mini {
    margin-left: 5px;
  }
  .warning {
    color: #dc143c;
  }
</style>