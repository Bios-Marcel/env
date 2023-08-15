module module_a

go 1.20

require github.com/caarlos0/env/v9 v9.0.0 // indirect

require module_b v0.0.0

replace module_b => ../module_b
