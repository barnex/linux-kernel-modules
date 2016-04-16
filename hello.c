#include <linux/init.h>
#include <linux/module.h>

MODULE_LICENSE("GPL");

#define NAME "hello"

static int __init mod_init(void) {
	printk(KERN_INFO "%s: init\n", NAME);
	return 0;
}

static void __exit mod_exit(void) {
	printk(KERN_INFO "%s: exit\n", NAME);
}

module_init(mod_init);
module_exit(mod_exit);
