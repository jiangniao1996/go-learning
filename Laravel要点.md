## 服务容器













## 生命周期

- 加载composer扩展文件

```php
require __DIR__ . '/../vendor/autoload.php';
```

- 创建容器，并在容器里绑定HTTP服务，命令行服务，异常处理服务

```php
$app = require __DIR__ . '/../bootstrap/app.php';
```

- 将处理HTTP的服务提供者解析出来，让该服务处理请求，并将请求返回

```php
$kernel = $app->app(Illuminate\Contracts\Http\Kernel::class);

$response = $kernel->handle(
	$request = Illuminate\Http\Request::capture()
);

$response->send();
```









## 门面



```php
Cache::get('name');

$app->make('cache')->get('');
```











## 版本差别













## 要点案例——商城秒杀

