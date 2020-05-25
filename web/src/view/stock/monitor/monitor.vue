<template>
  <div>
    <div class="search-term">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
        <el-row>
          <el-col :span="5">
            <el-form-item label="名称">
              <el-input placeholder="平安银行" v-model="searchInfo.name"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="5">
            <el-form-item label="编号">
              <el-input placeholder="SZ0000001"  v-model="searchInfo.symbol"></el-input>
            </el-form-item>
          </el-col>
        </el-row>
        <el-row>
          <el-col :span="5">
            <el-form-item label="市值">
              <el-input-number placeholder="最小" v-model="searchInfo.marketCapitalMin" :controls="false"></el-input-number>
            </el-form-item>
          </el-col>
          <el-col :span="5">
            <el-form-item label="">
              <el-input-number placeholder="最大"  v-model="searchInfo.marketCapitalMax" :controls="false"></el-input-number>
            </el-form-item>
          </el-col>
        </el-row>
        <el-row>
          <el-col :span="5">
            <el-form-item label="涨幅">
              <el-input-number placeholder="最小"  v-model="searchInfo.percentMin" :controls="false"></el-input-number>
            </el-form-item>
          </el-col>
          <el-col :span="5">
            <el-form-item label="">
              <el-input-number placeholder="最大"  v-model="searchInfo.percentMax" :controls="false"></el-input-number>
            </el-form-item>
          </el-col>
        </el-row>
        <el-row>
          <el-col :span="5">
            <el-form-item label="量比">
              <el-input-number placeholder="最小"  v-model="searchInfo.volume_ratio_min" :controls="false"></el-input-number>
            </el-form-item>
          </el-col>
          <el-col :span="5">
            <el-form-item label="">
              <el-input-number placeholder="最大"  v-model="searchInfo.volume_ratio_max" :controls="false"></el-input-number>
            </el-form-item>
          </el-col>
        </el-row>
        <el-form-item>
          <el-button @click="onSubmit" type="primary">查询</el-button>
        </el-form-item>
        <el-form-item>
          <el-button @click="openDialog('addStockMonitor')" type="primary">新增api</el-button>
        </el-form-item>
      </el-form>
    </div>
    <el-table :data="tableData" border stripe :default-sort = "{prop: ['market_capital','percent','volume_ratio','high','limit_down','chg','low','volume','amount','open','last_close'], order: 'descending'}">
      <el-table-column label="名称" min-width="150" prop="symbol"></el-table-column>
      <el-table-column label="市值" min-width="130" prop="market_capital" sortable ></el-table-column>

      <el-table-column label="涨幅" min-width="80" prop="percent" sortable></el-table-column>
      <el-table-column label="量比" min-width="80" prop="volume_ratio" sortable></el-table-column>

      <!--      <el-table-column label="涨停" min-width="80" prop="high" sortable></el-table-column>
            <el-table-column label="跌停" min-width="80" prop="limit_down" sortable></el-table-column>-->

      <!--      <el-table-column label="最高" min-width="80" prop="chg" sortable></el-table-column>
            <el-table-column label="最低" min-width="80" prop="low" sortable></el-table-column>-->

      <el-table-column label="成交量" min-width="120" prop="volume" sortable></el-table-column>
      <el-table-column label="成交额" min-width="120" prop="amount" sortable></el-table-column>

      <!--      <el-table-column label="今开" min-width="80" prop="open" sortable></el-table-column>
            <el-table-column label="昨收" min-width="80" prop="last_close" sortable></el-table-column>-->
    </el-table>

    <el-dialog :before-close="closeDialog" :title="dialogTitle" :visible.sync="dialogFormVisible">
      <el-form :inline="true" :model="form" :rules="rules" label-width="80px" ref="apiForm">
        <el-form-item label="路径" prop="path">
          <el-input autocomplete="off" v-model="form.path"></el-input>
        </el-form-item>
        <el-form-item label="请求" prop="method">
          <el-select placeholder="请选择" v-model="form.method">
            <el-option
                    :key="item.value"
                    :label="`${item.label}(${item.value})`"
                    :value="item.value"
                    v-for="item in methodOptions"
            ></el-option>
          </el-select>
        </el-form-item>
        <el-form-item label="api分组" prop="apiGroup">
          <el-input autocomplete="off" v-model="form.apiGroup"></el-input>
        </el-form-item>
        <el-form-item label="api简介" prop="description">
          <el-input autocomplete="off" v-model="form.description"></el-input>
        </el-form-item>
      </el-form>
      <div class="warning">新增监控需要在角色管理内配置权限才可使用</div>
      <div class="dialog-footer" slot="footer">
        <el-button @click="closeDialog">取 消</el-button>
        <el-button @click="enterDialog" type="primary">确 定</el-button>
      </div>
    </el-dialog>

  </div>
</template>


<script>
  import {
    getMonitorOne,
    getMonitorList,
    updateMonitor,
    delMonitor,
    addMonitor,
  } from '@/api/stockMonitor'
  import infoList from '@/components/mixins/infoList'
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
    name: 'Top',
    mixins: [infoList],
    data() {
      return {
        dialogTitle: '新增监控',
        dialogFormVisible: false,
        listApi: getMonitorList,
        searchInfo: {
          name:undefined,
          symbol:undefined,
          marketCapitalMin: undefined,
          marketCapitalMax: undefined,
          percentMin: undefined,
          percentMax: undefined,
          volume_ratio_min: undefined,
          volume_ratio_max: undefined
        },
        methodOptions: methodOptions,
        type: '',
      }
    },
    methods: {
      //搜索
      onSubmit() {
        this.getTableData()
        setInterval(()=>{
          this.getTableData()
        },10000)
      },
      openDialog(type) {
        switch (type) {
          case 'add':
            this.dialogTitlethis = '新增监控'
            break
          case 'edit':
            this.dialogTitlethis = '编辑监控'
            break
          default:
            break
        }
        this.type = type
        this.dialogFormVisible = true
      },
      async enterDialog() {
        this.$refs.apiForm.validate(async valid => {
          if (valid) {
            switch (this.type) {
              case 'addApi':
              {
                const res = await addMonitor(this.form)
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
                const res = await updateMonitor(this.form)
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
      },
      closeDialog() {
        this.initForm()
        this.dialogFormVisible = false
      },
      async editApi(row) {
        const res = await getMonitorOne({ id: row.ID })
        this.form = res.data.api
        this.openDialog('edit')
      },
      async deleteApi(row) {
        this.$confirm('此操作将永久删除所有角色下该菜单, 是否继续?', '提示', {
          confirmButtonText: '确定',
          cancelButtonText: '取消',
          type: 'warning'
        }).then(async () => {
            const res = await delMonitor(row)
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
  }
  .el-button {
    float: right;
  }

  .el-tag--mini {
    margin-left: 5px;
  }
  .warning {
    color: #dc143c;
  }
</style>