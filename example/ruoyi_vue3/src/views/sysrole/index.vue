<template>
   <div class="app-container">
      <el-form :model="queryParams" ref="queryRef" :inline="true" v-show="showSearch" label-width="68px">

      <el-form-item label="ID" prop="id">
      <el-input
         v-model="queryParams.id"
         placeholder="请输入ID"
         clearable
         style="width: 240px"
         @keyup.enter="handleQuery"
      />
    </el-form-item>
      <el-form-item label="RoleName" prop="roleName">
      <el-input
         v-model="queryParams.roleName"
         placeholder="请输入RoleName"
         clearable
         style="width: 240px"
         @keyup.enter="handleQuery"
      />
    </el-form-item>
      <el-form-item label="RoleNameEn" prop="roleNameEn">
      <el-input
         v-model="queryParams.roleNameEn"
         placeholder="请输入RoleNameEn"
         clearable
         style="width: 240px"
         @keyup.enter="handleQuery"
      />
    </el-form-item>
      <el-form-item label="Status" prop="status">
      <el-input
         v-model="queryParams.status"
         placeholder="请输入Status"
         clearable
         style="width: 240px"
         @keyup.enter="handleQuery"
      />
    </el-form-item>
      <el-form-item label="Priority" prop="priority">
      <el-input
         v-model="queryParams.priority"
         placeholder="请输入Priority"
         clearable
         style="width: 240px"
         @keyup.enter="handleQuery"
      />
    </el-form-item>
      <el-form-item label="Comment" prop="comment">
      <el-input
         v-model="queryParams.comment"
         placeholder="请输入Comment"
         clearable
         style="width: 240px"
         @keyup.enter="handleQuery"
      />
    </el-form-item>
      <el-form-item label="Deleted" prop="deleted">
      <el-input
         v-model="queryParams.deleted"
         placeholder="请输入Deleted"
         clearable
         style="width: 240px"
         @keyup.enter="handleQuery"
      />
    </el-form-item>
      <el-form-item label="LastmodifiedBy" prop="lastmodifiedBy">
      <el-input
         v-model="queryParams.lastmodifiedBy"
         placeholder="请输入LastmodifiedBy"
         clearable
         style="width: 240px"
         @keyup.enter="handleQuery"
      />
    </el-form-item>
    <el-form-item label="Lastmodified" style="width: 308px">
    <el-date-picker
       v-model="dateRange"
       value-format="YYYY-MM-DD"
       type="daterange"
       range-separator="-"
       start-placeholder="开始时间"
       end-placeholder="结束时间"
    ></el-date-picker>
    </el-form-item>
      <el-form-item label="CreatedBy" prop="createdBy">
      <el-input
         v-model="queryParams.createdBy"
         placeholder="请输入CreatedBy"
         clearable
         style="width: 240px"
         @keyup.enter="handleQuery"
      />
    </el-form-item>
    <el-form-item label="Created" style="width: 308px">
    <el-date-picker
       v-model="dateRange"
       value-format="YYYY-MM-DD"
       type="daterange"
       range-separator="-"
       start-placeholder="开始时间"
       end-placeholder="结束时间"
    ></el-date-picker>
    </el-form-item>
    <el-form-item label="CreatedAt" style="width: 308px">
    <el-date-picker
       v-model="dateRange"
       value-format="YYYY-MM-DD"
       type="daterange"
       range-separator="-"
       start-placeholder="开始时间"
       end-placeholder="结束时间"
    ></el-date-picker>
    </el-form-item>
      <el-form-item label="UpdatedBy" prop="updatedBy">
      <el-input
         v-model="queryParams.updatedBy"
         placeholder="请输入UpdatedBy"
         clearable
         style="width: 240px"
         @keyup.enter="handleQuery"
      />
    </el-form-item>
    <el-form-item label="UpdatedAt" style="width: 308px">
    <el-date-picker
       v-model="dateRange"
       value-format="YYYY-MM-DD"
       type="daterange"
       range-separator="-"
       start-placeholder="开始时间"
       end-placeholder="结束时间"
    ></el-date-picker>
    </el-form-item>


         <el-form-item>
            <el-button type="primary" icon="Search" @click="handleQuery">搜索</el-button>
            <el-button icon="Refresh" @click="resetQuery">重置</el-button>
         </el-form-item>
      </el-form>

      <el-row :gutter="10" class="mb8">
         <el-col :span="1.5">
            <el-button
               type="primary"
               plain
               icon="Plus"
               @click="handleAdd"
               v-hasPermi="['system:sysrole:add']"
            >新增</el-button>
         </el-col>
         <el-col :span="1.5">
            <el-button
               type="success"
               plain
               icon="Edit"
               :disabled="single"
               @click="handleUpdate"
               v-hasPermi="['system:sysrole:edit']"
            >修改</el-button>
         </el-col>
         <el-col :span="1.5">
            <el-button
               type="danger"
               plain
               icon="Delete"
               :disabled="multiple"
               @click="handleDelete"
               v-hasPermi="['system:sysrole:remove']"
            >删除</el-button>
         </el-col>
         <!--
         <el-col :span="1.5">
            <el-button
               type="warning"
               plain
               icon="Download"
               @click="handleExport"
               v-hasPermi="['system:sysrole:export']"
            >导出</el-button>
         </el-col>
         
         <el-col :span="1.5">
            <el-button
               type="danger"
               plain
               icon="Refresh"
               @click="handleRefreshCache"
               v-hasPermi="['system:sysrole:remove']"
            >刷新缓存</el-button>
         </el-col>
         -->
         <right-toolbar v-model:showSearch="showSearch" @queryTable="getList"></right-toolbar>
      </el-row>

      <el-table v-loading="loading" :data="typeList" @selection-change="handleSelectionChange">
         <el-table-column type="selection" width="55" align="center" />

      <el-table-column label="ID" align="center" prop="id" />
      <el-table-column label="RoleName" align="center" prop="roleName" />
      <el-table-column label="RoleNameEn" align="center" prop="roleNameEn" />
      <el-table-column label="Status" align="center" prop="status" />
      <el-table-column label="Priority" align="center" prop="priority" />
      <el-table-column label="Comment" align="center" prop="comment" />
      <el-table-column label="Deleted" align="center" prop="deleted" />
      <el-table-column label="LastmodifiedBy" align="center" prop="lastmodifiedBy" />
    <el-table-column label="Lastmodified" align="center" prop="lastmodified" width="180">
      <template #default="scope">
         <span>{{ parseTime(scope.row.lastmodified) }}</span>
      </template>
      </el-table-column>
      <el-table-column label="CreatedBy" align="center" prop="createdBy" />
    <el-table-column label="Created" align="center" prop="created" width="180">
      <template #default="scope">
         <span>{{ parseTime(scope.row.created) }}</span>
      </template>
      </el-table-column>
    <el-table-column label="CreatedAt" align="center" prop="createdAt" width="180">
      <template #default="scope">
         <span>{{ parseTime(scope.row.createdAt) }}</span>
      </template>
      </el-table-column>
      <el-table-column label="UpdatedBy" align="center" prop="updatedBy" />
    <el-table-column label="UpdatedAt" align="center" prop="updatedAt" width="180">
      <template #default="scope">
         <span>{{ parseTime(scope.row.updatedAt) }}</span>
      </template>
      </el-table-column>


         <el-table-column label="操作" align="center" class-name="small-padding fixed-width">
            <template #default="scope">
               <el-button
                  type="text"
                  icon="Edit"
                  @click="handleUpdate(scope.row)"
                  v-hasPermi="['system:sysrole:edit']"
               >修改</el-button>
               <el-button
                  type="text"
                  icon="Delete"
                  @click="handleDelete(scope.row)"
                  v-hasPermi="['system:sysrole:remove']"
               >删除</el-button>
            </template>
         </el-table-column>
      </el-table>

      <pagination
         v-show="total > 0"
         :total="total"
         v-model:page="queryParams.page"
         v-model:limit="queryParams.pageSize"
         @pagination="getList"
      />

      <!-- 添加或修改参数配置对话框 -->
      <el-dialog :title="title" v-model="open" width="500px" append-to-body>
         <el-form ref="sysRoleRef" :model="form" :rules="rules" label-width="80px">

        <el-form-item label="RoleName" prop="roleName">
          <el-input v-model="form.roleName" placeholder="请输入RoleName" />
        </el-form-item>
        <el-form-item label="RoleNameEn" prop="roleNameEn">
          <el-input v-model="form.roleNameEn" placeholder="请输入RoleNameEn" />
        </el-form-item>
        <el-form-item label="Status" prop="status">
          <el-input v-model="form.status" placeholder="请输入Status" />
        </el-form-item>
        <el-form-item label="Priority" prop="priority">
          <el-input v-model="form.priority" placeholder="请输入Priority" />
        </el-form-item>
        <el-form-item label="Comment" prop="comment">
          <el-input v-model="form.comment" placeholder="请输入Comment" />
        </el-form-item>
        <el-form-item label="Deleted" prop="deleted">
          <el-input v-model="form.deleted" placeholder="请输入Deleted" />
        </el-form-item>
        <el-form-item label="LastmodifiedBy" prop="lastmodifiedBy">
          <el-input v-model="form.lastmodifiedBy" placeholder="请输入LastmodifiedBy" />
        </el-form-item>
        <el-form-item label="CreatedBy" prop="createdBy">
          <el-input v-model="form.createdBy" placeholder="请输入CreatedBy" />
        </el-form-item>
        <el-form-item label="UpdatedBy" prop="updatedBy">
          <el-input v-model="form.updatedBy" placeholder="请输入UpdatedBy" />
        </el-form-item>


         </el-form>
         <template #footer>
            <div class="dialog-footer">
               <el-button type="primary" @click="submitForm">确 定</el-button>
               <el-button @click="cancel">取 消</el-button>
            </div>
         </template>
      </el-dialog>
   </div>
</template>

<script setup name="SysRole">
import { listSysRole, getSysRole, delSysRole, addSysRole, updateSysRole } from "@/api/sysrole/sysrole";

const { proxy } = getCurrentInstance();

const typeList = ref([]);
const open = ref(false);
const loading = ref(true);
const showSearch = ref(true);
const ids = ref([]);
const single = ref(true);
const multiple = ref(true);
const total = ref(0);
const title = ref("");
const dateRange = ref([]);

const data = reactive({
  form: {},
  queryParams: {
    page: 1,
    pageSize: 10,
            id: undefined,
        roleName: undefined,
        roleNameEn: undefined,
        status: undefined,
        priority: undefined,
        comment: undefined,
        deleted: undefined,
        lastmodifiedBy: undefined,
        lastmodified: undefined,
        createdBy: undefined,
        created: undefined,
        createdAt: undefined,
        updatedBy: undefined,
        updatedAt: undefined,

   //  dictName: undefined,
   //  dictType: undefined,
   //  status: undefined
  },
//   rules: {
//     dictName: [{ required: true, message: "字典名称不能为空", trigger: "blur" }],
//     dictType: [{ required: true, message: "SysRole不能为空", trigger: "blur" }]
//   },
});

const { queryParams, form, rules } = toRefs(data);

/** 查询列表 */
function getList() {
  loading.value = true;
  listSysRole(proxy.addDateRange(queryParams.value, dateRange.value)).then(response => {
    typeList.value = response.data;
    total.value = response.total;
    loading.value = false;
  });
}
/** 取消按钮 */
function cancel() {
  open.value = false;
  reset();
}
/** 表单重置 */
function reset() {
  form.value = {
   //  id: undefined,
   //  dictName: undefined,
   //  dictType: undefined,
           id: undefined,
        roleName: undefined,
        roleNameEn: undefined,
        status: undefined,
        priority: undefined,
        comment: undefined,
        deleted: undefined,
        lastmodifiedBy: undefined,
        lastmodified: undefined,
        createdBy: undefined,
        created: undefined,
        createdAt: undefined,
        updatedBy: undefined,
        updatedAt: undefined,

    status: "0",
    remark: undefined
  };
  proxy.resetForm("sysRoleRef");
}
/** 搜索按钮操作 */
function handleQuery() {
  queryParams.value.page = 1;
  getList();
}
/** 重置按钮操作 */
function resetQuery() {
  dateRange.value = [];
  proxy.resetForm("queryRef");
  handleQuery();
}
/** 新增按钮操作 */
function handleAdd() {
  reset();
  open.value = true;
  title.value = "添加SysRole";
}
/** 多选框选中数据 */
function handleSelectionChange(selection) {
  ids.value = selection.map(item => item.id);
  single.value = selection.length != 1;
  multiple.value = !selection.length;
}
/** 修改按钮操作 */
function handleUpdate(row) {
  reset();
  const id = row.id || ids.value;
  getSysRole(id).then(response => {
    form.value = response.data;
    open.value = true;
    title.value = "修改SysRole";
  });
}
/** 提交按钮 */
function submitForm() {
  proxy.$refs["sysRoleRef"].validate(valid => {
    if (valid) {
      if (form.value.id != undefined) {
        updateSysRole(form.value).then(response => {
          proxy.$modal.msgSuccess("修改成功");
          open.value = false;
          getList();
        });
      } else {
        addSysRole(form.value).then(response => {
          proxy.$modal.msgSuccess("新增成功");
          open.value = false;
          getList();
        });
      }
    }
  });
}
/** 删除按钮操作 */
function handleDelete(row) {
  const ids = row.id || ids.value;
  proxy.$modal.confirm('是否确认删除ID为"' + ids + '"的数据项？').then(function() {
    return delSysRole(ids);
  }).then(() => {
    getList();
    proxy.$modal.msgSuccess("删除成功");
  }).catch(() => {});
}

// /** 导出按钮操作 */
// function handleExport() {
//   proxy.download("system/dict/type/export", {
//     ...queryParams.value
//   }, `dict_${new Date().getTime()}.xlsx`);
// }

/** 刷新缓存按钮操作 */
// function handleRefreshCache() {
//   refreshCache().then(() => {
//     proxy.$modal.msgSuccess("刷新成功");
//   });
// }

getList();
</script>
