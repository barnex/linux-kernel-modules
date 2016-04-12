//#include <asm/uaccess.h>
//#include <linux/device.h>
//#include <linux/fs.h>
//#include <linux/init.h>
//#include <linux/kernel.h>
#include <linux/module.h>

#define DEVICE_NAME "chardev" // device will be at /dev/chardev
#define CLASS_NAME  "arne"

MODULE_LICENSE("GPL");
MODULE_AUTHOR("barnex");
MODULE_DESCRIPTION("my first character device");
MODULE_VERSION("0.0");

//static int majorNumber; // device number
//static char message[256];
//static int len_message;
//static int countOpens;
//static struct class* drvClass = NULL;
//static struct device* drvDevice = NULL;

//static int dev_open(struct inode *, struct file *);
//static ssize_t dev_read(struct file *, char *, size_t, loff_t *);
//static int dev_release(struct inode *, struct file *);
//static ssize_t dev_write(struct file *, const char *, size_t, loff_t *);
//
//static struct file_operations fops = {
//	.open = dev_open,
//	.read = dev_read,
//	.release = dev_release,
//	.write= dev_write,
//};

static int __init chardev_init(void){
	printk(KERN_INFO "Arne's character device loading\n");
	return 0;
}

static void __exit chardev_exit(void){
	printk(KERN_INFO "Arne's character device unloading\n");
}

module_init(chardev_init);
module_exit(chardev_exit);
