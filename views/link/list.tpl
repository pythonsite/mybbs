<div class="row">
  <div class="col-md-12">
    <div class="panel panel-default">
      <div class="panel-heading">
        友链管理
        <a href="/link/add" class="pull-right">添加友链</a>
      </div>
      <div class="table-responsive">
        <table class="table table-striped table-responsive">
          <tbody>
          {{range .FriendLinks}}
          <tr>
            <td>{{.Id}}</td>
            <td>{{.LinkName}}</td>
            <td>{{.LinkUrl}}</td>
            <td>
              <a href="/link/edit/{{.Id}}" class="btn btn-xs btn-warning">编辑</a>
              <a href="javascript:if(confirm('确认删除吗?')) location.href='/link/delete/{{.Id}}'" class="btn btn-xs btn-danger">删除</a>
            </td>
          </tr>
          {{end}}
          </tbody>
        </table>
      </div>
    </div>
  </div>
</div>