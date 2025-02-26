# Benchmark de Geração de UUID sem Bibliotecas

Este repositório contém implementações de geração de UUID v4 em diferentes linguagens de programação sem o uso de bibliotecas externas. O objetivo é comparar a eficiência de cada linguagem em termos de tempo de execução.



📌 Tecnologias Testadas

* JavaScript (Executado no Node.js)

* Go (Golang)

* Python

* PHP

🚀 Como Executar os Testes

JavaScript (Node.js)

Execute o seguinte comando para rodar em JavaScript:

```bash
node generate-uuid.js
```

Go (Golang)

Compile e execute o código Go com:

```bash
go run generate-uuid.go
``` 

Python

Execute em Python com:

```bash
python generate-uuid.py
```

PHP

Execute em PHP com:

```bash
php generate-uuid.php
```

⏱️ Resultados do Benchmark

Os tempos abaixo representam a execução de 1.000.000 de UUIDs gerados em cada linguagem.

| Linguagem | Tempo de Execução |
|-----------|-------------------|
| JavaScript | 1.933s |
| Go | 1.038s |
| Python | 0.880s |
| PHP | 0.549s |

📊 Análise de Desempenho

* PHP foi a mais rápida, apesar de ser interpretada, devido ao uso otimizado de mt_rand().

* Python apresentou um bom desempenho, ficando atrás apenas do PHP.

* Go teve um tempo intermediário, mesmo sendo uma linguagem compilada, pois crypto/rand é mais seguro, mas mais lento que mt_rand().

* JavaScript foi a mais lenta, pois Math.random() não é tão eficiente para geração de números aleatórios de alta entropia.


🛠️ Implementações

JavaScript

```javascript
function generateUUID() {
  return 'xxxxxxxx-xxxx-4xxx-yxxx-xxxxxxxxxxxx'.replace(/[xy]/g, function (c) {
    var r = (Math.random() * 16) | 0,
      v = c == 'x' ? r : (r & 0x3) | 0x8;
    return v.toString(16);
  });
}
```

Go

```go
func generateUUID() string {
	uuid := make([]byte, 16)
	_, err := rand.Read(uuid)
	if err != nil {
		return ""
	}

	uuid[6] = (uuid[6] & 0x0F) | 0x40
	uuid[8] = (uuid[8] & 0x3F) | 0x80

	return fmt.Sprintf("%08x-%04x-%04x-%04x-%012x",
		uuid[0:4], uuid[4:6], uuid[6:8], uuid[8:10], uuid[10:16])
}
```

Python

```python
def manual_uuid():
    return "%08x-%04x-%04x-%04x-%012x" % (
        random.getrandbits(32),
        random.getrandbits(16),
        (random.getrandbits(16) & 0x0fff) | 0x4000,
        (random.getrandbits(16) & 0x3fff) | 0x8000,
        random.getrandbits(48),
    )
```

PHP

```php
function generateUuid() {
  return sprintf(
    '%04x%04x-%04x-%04x-%04x-%04x%04x%04x',
    mt_rand(0, 0xffff),
    mt_rand(0, 0xffff),
    mt_rand(0, 0xffff),
    mt_rand(0, 0x0fff) | 0x4000,
    mt_rand(0, 0x3fff) | 0x8000,
    mt_rand(0, 0xffff),
    mt_rand(0, 0xffff),
    mt_rand(0, 0xffff)
  );
}
```

📌 Conclusão

Se você precisa de performance máxima para geração de UUIDs:

* Use PHP para rapidez, mas cuidado com segurança para uso criptográfico.

* Use Python para um equilíbrio entre velocidade e segurança.

* Use Go se precisar de um UUID mais seguro com crypto/rand.

* Evite JavaScript para geração massiva de UUIDs, pois é a opção mais lenta.

🔗 Contribuições e sugestões são bem-vindas! 🚀