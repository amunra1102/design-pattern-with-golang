package adapter

type Client struct {}

func (c *Client) InsertSquareUSBInComputer(computer Computer) {
	computer.insertInSquarePort()
}
