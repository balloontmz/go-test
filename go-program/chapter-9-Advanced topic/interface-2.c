#include <stdio.h>
#include <stdlib.h>

// ----------------------------------------------------------------
typedef struct _TypeInfo {
    // 用于运行时取得类型信息，比如反射机制
} TypeInfo;

typedef struct _InterfaceInfo{
    // 用于运行时取得接口信息
} InterfaceInfo;

// ----------------------------------------------------------------
typedef struct _IReadWriterTbl {
    InterfaceInfo* inter;
    TypeInfo* type;
    int (*Read) (void* this, char* buf, int cb);
    int (*Write) (void* this, char* buf, int cb);
} IReadWriterTbl;

typedef struct _IReadWriter {
    IReadWriterTbl* tab;
    void* data;
} IReadWriter;

InterfaceInfo g_InterfaceInfo_IReadWriter = {
    //
};

// ----------------------------------------------------------------

typedef struct _A {
    int a;
} A;

int A_Read(A* this, char* buf, int cb) {
    printf("A_Read: %d\n", this->a);
    return cb;
}

int A_Write(A* this, char* buf, int cb) {
    printf("A_Write: %d\n", this->a);
    return cb;
}

TypeInfo g_TypeInfo_A = {
    //
};

A* NewA(int params) {
    printf("NewA: %d\n", params);
    // 动态内存分配，用于申请一块连续的指定大小的内存块区域以void*类型返回分配的内存区域地址
    A* this = (A*)malloc(sizeof(A));
    this->a = params;
    return this;
}

// ----------------------------------------------------------------
typedef struct _B {
    A base;
} B;

int B_Write(B* this, char* buf, int cb) {
    printf("B_Write: %d\n", this->base.a);
    return cb;
}

void B_Foo(B* this) {
    printf("B_Foo: %d\n", this->base.a);
}

TypeInfo g_TypeInfo_B = {

};

B* NewB(int params) {
    printf("NewB: %d\n", params);
    // 动态内存分配，用于申请一块连续的指定大小的内存块区域以void*类型返回分配的内存区域地址
    B* this = (B*)malloc(sizeof(A));
    this->base.a = params;
    return this;
}

// ----------------------------------------------------------------

IReadWriterTbl g_Ital_IReadWriter_B = {
    &g_InterfaceInfo_IReadWriter,
    &g_TypeInfo_B,
    // c 的函数原型？？？
    // 对 A 的函数传入 B 的对象？
    //TODO: 此处暂时理解为 指向 B 的指针其实也是指向 A 的指针，因为 B 只有 A 一个成员;
    (int (*)(void* this, char* buf, int cb))A_Read,
    (int (*)(void* this, char* buf, int cb))B_Write,
};

int main() {
    B* unamed = NewB(8);
    IReadWriter p = {
        &g_Ital_IReadWriter_B,
        unamed
    };

    p.tab->Read(p.data, NULL, 10);
    p.tab->Write(p.data, NULL, 10);
    return EXIT_SUCCESS;
}