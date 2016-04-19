#ifndef __HELLO_H_
#define __HELLO_H_

#define DEVICE_NAME "hello" // device will be at /dev/hello
#define CLASS_NAME  "hellodevice"

static int major;
static struct class* dev_class;
static struct device* dev_device;

static char message[256];
static int len_message;

static int __init mod_init(void);
static void mod_exit(void);

static int dev_open(struct inode *inodep, struct file *filep);
static ssize_t dev_read(struct file *filep, char *buffer, size_t len, loff_t *offset);
static int dev_release(struct inode *inodep, struct file *filep);

#endif
