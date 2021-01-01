echo "コード生成開始"
goa gen calc/design
goa example calc/design
echo "コード生成完了"

echo "ビルド開始"
go build ./cmd/calc && go build ./cmd/calc-cli
echo "ビルド完了"