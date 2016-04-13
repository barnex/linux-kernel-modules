//#include <asm/uaccess.h>
#include <linux/device.h>
#include <linux/fs.h>
//#include <linux/init.h>
//#include <linux/kernel.h>
#include <linux/module.h>

#define DEVICE_NAME "chardev" // device will be at /dev/chardev
#define CLASS_NAME  "arne"

MODULE_LICENSE("GPL");
MODULE_AUTHOR("barnex");
MODULE_DESCRIPTION("my first character device");
MODULE_VERSION("0.0");

static int majorNumber; // device number
static char message[256];
static int len_message;
static int countOpens;
static struct class* myClass = NULL;
static struct device* myDevice = NULL;

static int dev_open(struct inode *inodep, struct file *filep){
	return 0;
}

static ssize_t dev_read(struct file *filep, char *buffer, size_t len, loff_t *offset){
	return 0;
}

static int dev_release(struct inode *inodep, struct file *filep){
	return 0;
}

static ssize_t dev_write(struct file *filep, const char *buffer, size_t len, loff_t *offset){
	return 0;
}


static struct file_operations fops = {
	.open = dev_open,
	.read = dev_read,
	.release = dev_release,
	.write= dev_write,
};

static int __init chardev_init(void){
	printk(KERN_INFO "Arne's character device: loading\n");

	majorNumber = register_chrdev(0, DEVICE_NAME, &fops);
	if(majorNumber < 0){
		printk(KERN_ALERT "Arne's character device: register_chrdev failed\n");
		return majorNumber;	
	}
	printk(KERN_INFO "Arne's character device: register_chrdev: %d\n", majorNumber);

	myClass = class_create(THIS_MODULE, CLASS_NAME);
	if (IS_ERR(myClass)){
		class_destroy(myClass);
		unregister_chrdev(majorNumber, DEVICE_NAME);	
		printk(KERN_ALERT "Arne's character device: device_create failed\n");
		return PTR_ERR(myClass);
	}
	printk(KERN_INFO "Arne's character device: class_create: OK\n");



	return 0;
}

static void __exit chardev_exit(void){
	printk(KERN_INFO "Arne's character device: unloading\n");
}

module_init(chardev_init);
module_exit(chardev_exit);
