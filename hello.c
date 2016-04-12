#include <linux/module.h>
#include <linux/kernel.h>
#include <linux/init.h>

MODULE_LICENSE("GPL");
MODULE_AUTHOR("barnex");
MODULE_DESCRIPTION("Hello, kernel module!");

static int __init hello_init(void) {
    printk(KERN_INFO "Hello, kernel module!\n");
    return 0;
}

static void __exit hello_cleanup(void) {
    printk(KERN_INFO "Bye bye, kernel module.\n");
}

module_init(hello_init);
module_exit(hello_cleanup);
