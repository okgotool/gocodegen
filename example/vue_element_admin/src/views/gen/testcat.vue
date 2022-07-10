<template>
  <div class="app-container">
    <el-row>
      <el-button type="primary" @click="handleOpenCreateDialog">创建</el-button>
    </el-row>
    <el-dialog
      title="创建"
      :visible.sync="dialogVisible"
      width="50%"
      :before-close="handleClose"
    >
      <el-form ref="form" :model="form" label-width="32%">
        <el-form-item label="CatName">
          <el-input v-model="form.cat_name"></el-input>
        </el-form-item>
        <el-form-item label="CreatedBy">
          <el-input v-model="form.created_by"></el-input>
        </el-form-item>
			<el-form-item label="CreatedAt">
			<el-date-picker
			  type="date"
			  placeholder="选择日期"
			  v-model="form.created_at"
			  style="width: 100%"
			></el-date-picker>
		  </el-form-item>
        <el-form-item label="UpdatedBy">
          <el-input v-model="form.updated_by"></el-input>
        </el-form-item>
			<el-form-item label="UpdatedAt">
			<el-date-picker
			  type="date"
			  placeholder="选择日期"
			  v-model="form.updated_at"
			  style="width: 100%"
			></el-date-picker>
		  </el-form-item>

        <el-form-item>
          <el-button type="primary" @click="handleCreate">保存</el-button>
          <el-button @click="onCancel">取消</el-button>
        </el-form-item>
      </el-form>
    </el-dialog>

    <el-dialog
      title="编辑"
      :visible.sync="editDialogVisible"
      width="50%"
      :before-close="handleClose"
    >
      <el-form ref="form" :model="form" label-width="32%">
        <el-form-item label="CatName">
          <el-input v-model="form.cat_name"></el-input>
        </el-form-item>
        <el-form-item label="CreatedBy">
          <el-input v-model="form.created_by"></el-input>
        </el-form-item>
			<el-form-item label="CreatedAt">
			<el-date-picker
			  type="date"
			  placeholder="选择日期"
			  v-model="form.created_at"
			  style="width: 100%"
			></el-date-picker>
		  </el-form-item>
        <el-form-item label="UpdatedBy">
          <el-input v-model="form.updated_by"></el-input>
        </el-form-item>
			<el-form-item label="UpdatedAt">
			<el-date-picker
			  type="date"
			  placeholder="选择日期"
			  v-model="form.updated_at"
			  style="width: 100%"
			></el-date-picker>
		  </el-form-item>

        <el-form-item>
          <el-button type="primary" @click="handleEditRow">保存</el-button>
          <el-button @click="onEditCancel">取消</el-button>
        </el-form-item>
      </el-form>
    </el-dialog>
     
    <el-table
      v-loading="listLoading"
      :data="list"
      height="500"
      border
      fit
      highlight-current-row
      style="width: 100%"
    >    
      <el-table-column align="left" sortable label="ID">
        <template slot-scope="{ row }">
          <span>{{ row.id }}</span>
        </template>
      </el-table-column>

      <el-table-column align="left" sortable label="CatName">
        <template slot-scope="{ row }">
          <span>{{ row.cat_name }}</span>
        </template>
      </el-table-column>

      <el-table-column align="left" sortable label="CreatedBy">
        <template slot-scope="{ row }">
          <span>{{ row.created_by }}</span>
        </template>
      </el-table-column>

      <el-table-column align="left" sortable label="CreatedAt">
        <template slot-scope="{ row }">
          <span>{{ row.created_at }}</span>
        </template>
      </el-table-column>

      <el-table-column align="left" sortable label="UpdatedBy">
        <template slot-scope="{ row }">
          <span>{{ row.updated_by }}</span>
        </template>
      </el-table-column>

      <el-table-column align="left" sortable label="UpdatedAt">
        <template slot-scope="{ row }">
          <span>{{ row.updated_at }}</span>
        </template>
      </el-table-column>



      <el-table-column fixed="right" label="操作" width="100">
        <template slot-scope="scope">
          <el-button @click="handleOpenEditDialog(scope.row)" type="text" size="small">
            编辑
          </el-button>
          <el-button
            @click="handleRemoveRow(scope.row)"
            type="text"
            size="small"
          >
            删除
          </el-button>
        </template>
      </el-table-column>
    </el-table>
  </div>
</template>

<script>
import { queryTestCatAll,createTestCat,updateTestCat,removeTestCat } from "@/api/testCat";

export default {
  name: "EditTable",
  data() {
    return {
      list: null,
      listLoading: true,
      listQuery: {
        page: 1,
        pageSize: 10,
      },
      dialogVisible: false,
      editDialogVisible: false,
      form: {
        id: 0,
        cat_name: '',
        created_by: '',
        created_at: '',
        updated_by: '',
        updated_at: '',

      },
    };
  },
  created() {
    this.getList();
  },
  methods: {
    async getList() {
      this.listLoading = true;
      const { data } = await queryTestCatAll(this.listQuery);
      const items = data;
      this.list = items.map((v) => {
        v.originalTitle = v.title; 
        return v;
      });
      this.listLoading = false;
    },
    handleOpenCreateDialog() {
      this.form = {};
      this.dialogVisible = true;
    },    
    handleClose(done) {
      // this.$confirm("确认关闭？")
      //  .then((_) => {
          done();
      //  })
      //  .catch((_) => {});
    },
    handleCreate() {
      console.log("handleCreate!");
      this.form.id = 0;
      const { data } = createTestCat(this.form);
      this.dialogVisible = false;
      this.getList();
    },
    handleOpenEditDialog(row){
      this.editDialogVisible = true;
      this.form = row;
     },
    handleEditRow() {
      const { data } = updateTestCat(this.form);
      this.editDialogVisible = false;
      this.getList();
    },
    handleRemoveRow(row) {
      this.$confirm("确认删除？")
        .then((_) => {
          let deleteParam = { id: row.id };
          console.log("handleRemoveRow:", deleteParam);
          const { data } = removeTestCat(deleteParam);
          this.dialogVisible = false;
          this.getList();
        })
        .catch((_) => {});
    },
    onCancel() {
      this.dialogVisible = false;
    },
    onEditCancel() {
      this.editDialogVisible = false;
    },
  },
};
</script>
