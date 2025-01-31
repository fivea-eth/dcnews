package main

import (
	"encoding/json"
	"log"
	"os"

	"github.com/bwmarrin/discordgo"
)

type Configuration struct {
	Discord_bot_auth string
}

func discord_send_text(content string, dc_channel_id string) {

	// 打开文件
	config_file, _ := os.Open("config.json")

	// 关闭文件
	defer config_file.Close()

	//NewDecoder创建一个从file读取并解码json对象的*Decoder，解码器有自己的缓冲，并可能超前读取部分json数据。
	decoder := json.NewDecoder(config_file)

	conf := Configuration{}
	//Decode从输入流读取下一个json编码值并保存在v指向的值里
	decoder.Decode(&conf)

	discord, err := discordgo.New("Bot " + conf.Discord_bot_auth)
	if err != nil {
		log.Println("Failed to create Discord session: ", err)
		return
	}

	// 建立连接
	err = discord.Open()
	if err != nil {
		log.Println("Failed to open connection: ", err)
	}

	// 发送消息
	_, err = discord.ChannelMessageSend(dc_channel_id, content)
	if err != nil {
		log.Println("Error sending text: ", err)
	}
	discord.Close()
}

func discord_send_file(content string, name string, path string, dc_channel_id string) {

	// 打开文件
	config_file, _ := os.Open("config.json")

	// 关闭文件
	defer config_file.Close()

	//NewDecoder创建一个从file读取并解码json对象的*Decoder，解码器有自己的缓冲，并可能超前读取部分json数据。
	decoder := json.NewDecoder(config_file)

	conf := Configuration{}
	//Decode从输入流读取下一个json编码值并保存在v指向的值里
	decoder.Decode(&conf)

	discord, err := discordgo.New("Bot " + conf.Discord_bot_auth)
	if err != nil {
		log.Println("Failed to create Discord session: ", err)
		return
	}

	// 建立连接
	err = discord.Open()
	if err != nil {
		log.Println("Failed to open connection: ", err)
	}

	file, err := os.Open(path)
	if err != nil {
		log.Println("Error opening image file: ", err)
		return
	}
	// 发送消息
	_, err = discord.ChannelFileSendWithMessage(dc_channel_id, content, name, file)
	if err != nil {
		log.Println("Error sending image: ", err)
	}

	discord.Close()

}
