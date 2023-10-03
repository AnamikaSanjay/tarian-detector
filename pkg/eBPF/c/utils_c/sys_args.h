#ifndef __UTILS_C_SYS_ARGS_H__
#define __UTILS_C_SYS_ARGS_H__

#include "index.h"

stain int init_sys_ctx(sys_ctx_t *);
stain int read_sys_ctx_into(sys_ctx_t *, struct pt_regs *);

stain int init_sys_ctx(sys_ctx_t *dst) {
  if (dst == NULL)
    return NULL_POINTER_ERROR;

  (*dst)[0] = 0;
  (*dst)[1] = 0;
  (*dst)[2] = 0;
  (*dst)[3] = 0;
  (*dst)[4] = 0;
  (*dst)[5] = 0;
  (*dst)[6] = 0;

  return OK;
};

stain int read_sys_ctx_into(sys_ctx_t *dst, struct pt_regs *src) {
  if (dst == NULL || src == NULL)
    return NULL_POINTER_ERROR;

  int err = init_sys_ctx(dst);
  if (err != OK)
    return err;

  struct pt_regs *ctx2 = (struct pt_regs *)PT_REGS_PARM1_CORE(src);

  (*dst)[0] = PT_REGS_PARM1_CORE(ctx2);
  (*dst)[1] = PT_REGS_PARM2_CORE(ctx2);
  (*dst)[2] = PT_REGS_PARM3_CORE(ctx2);
  (*dst)[3] = PT_REGS_PARM4_CORE(ctx2);
  (*dst)[4] = PT_REGS_PARM5_CORE(ctx2);
  (*dst)[5] = PT_REGS_PARM6_CORE(ctx2);
  (*dst)[6] = PT_REGS_SYSCALL_CORE(ctx2);

  return OK;
}

#endif