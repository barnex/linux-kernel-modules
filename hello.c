#include <asm/uaccess.h>
#include <linux/device.h>
#include <linux/fs.h>
#include <linux/module.h>
#include "hello.h"

#define DEVICE_NAME "hello" // device will be at /dev/hello
#define CLASS_NAME  "hellodevice"

MODULE_LICENSE("GPL");
MODULE_AUTHOR("barnex");
MODULE_DESCRIPTION("Fun with Linux char devices.");
MODULE_VERSION("0.3");

static int major;
static struct class* dev_class = NULL;
static struct device* dev_device = NULL;

static char message[256] = "Hello from the linux kernel!\n";
static int len_message = 29;

static int dev_open(struct inode *inodep, struct file *filep) {
	return 0;
}

static ssize_t dev_read(struct file *filep, char *buffer, size_t len, loff_t *offset) {
	int n;
	int err;

	if (*offset >= len_message) {
		return 0; //EOF
	}

	n = len;
	if (n > len_message - *offset) {
		n = len_message - *offset;
	}

	err = copy_to_user(buffer, &message[*offset], n);
	if (err != 0) {
		printk(KERN_ERR "hellodev: dev_read failed: %d\n", err);
		return -EFAULT;
	}
	*offset += n;
	return n;
}

static int dev_release(struct inode *inodep, struct file *filep) {
	return 0;
}

static struct file_operations fops = {
	.owner = THIS_MODULE,
	.open = dev_open,
	.read = dev_read,
	.release = dev_release,
};

static int __init mod_init(void) {
	major = register_chrdev(0, DEVICE_NAME, &fops);
	if(major < 0) {
		printk(KERN_ALERT "hellodev: register_chrdev failed\n");
		mod_exit();
		return major;
	}
	printk(KERN_INFO "hellodev: register_chrdev: %d\n", major);

	dev_class = class_create(THIS_MODULE, CLASS_NAME);
	if (IS_ERR(dev_class)) {
		printk(KERN_ALERT "hellodev: class_create failed\n");
		mod_exit();
		return PTR_ERR(dev_class);
	}
	printk(KERN_INFO "hellodev: class_create: OK\n");

	dev_device = device_create(dev_class, NULL, MKDEV(major, 0), NULL, DEVICE_NAME);
	if (IS_ERR(dev_device)) {
		printk(KERN_ALERT "hellodev: device_create failed\n");
		mod_exit();
		return PTR_ERR(dev_device);
	}
	printk(KERN_INFO "hellodev: device_create: OK\n");

	return 0;
}

static void mod_exit(void) {
	if (dev_device != NULL) {
		printk(KERN_INFO "hellodev: device_destroy\n");
		device_destroy(dev_class, MKDEV(major, 0));
	}
	if (dev_class != NULL) {
		printk(KERN_INFO "hellodev: class_destroy\n");
		class_unregister(dev_class);
		class_destroy(dev_class);
	}
	if (major != 0) {
		printk(KERN_INFO "hellodev: unregister_chrdev\n");
		unregister_chrdev(major, DEVICE_NAME);
	}
}

module_init(mod_init);
module_exit(mod_exit);
