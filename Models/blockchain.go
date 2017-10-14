package Models

//Blockchain model
type Blockchain struct {
	Blocks []*Block
}

//AddNewBlock - add new block to the chain
func (bc *Blockchain) AddNewBlock() {
	if len(bc.Blocks) < 1 {

	} else {

	}
}
