package main

/*
#cgo CFLAGS: -I ./scpi-parser/libscpi/inc
#cgo LDFLAGS: ./scpi-parser/libscpi/dist/libscpi.a

#include <stdio.h>
#include "./scpi-parser/libscpi/inc/scpi/scpi.h"

typedef struct scpi_object_s{
	scpi_t 				ctx;
	scpi_interface_t	interf;
	scpi_error_t		errq[16];

	scpi_command_t		cmds[16];
	int					cmd_idx;
}scpi_object_s;


static size_t scpi_write(scpi_t* ctx, const char* data, size_t len){
	printf("scpi_write(), len=%zu\n", len);
	return len;
}

static int scpi_error_callback(scpi_t* ctx, int_fast16_t error){
	printf("scpi_error_callback(), err=%ld\n",  error);
	return 0;
}

static scpi_result_t scpi_flush(scpi_t * context) {
	printf("scpi_flush()\n");
	return SCPI_RES_OK;
}

static scpi_result_t scpi_reset(scpi_t * context) {
	printf("scpi_reset()\n");
	return SCPI_RES_OK;
}

static scpi_result_t scpi_control(scpi_t * context, scpi_ctrl_name_t ctrl, scpi_reg_val_t val) {
    if (SCPI_CTRL_SRQ == ctrl) {
        fprintf(stderr, "**SRQ: 0x%X (%d)\r\n", val, val);
    } else {
        fprintf(stderr, "**CTRL %02x: 0x%X (%d)\r\n", ctrl, val, val);
    }
    return SCPI_RES_OK;
}

// scpi handler
// must impl outside
scpi_result_t hdl_scpi(scpi_t * context) __attribute__((weak));
scpi_result_t hdl_scpi(scpi_t * context){
	printf("need impl, uid=%lu\n", SCPI_CmdTag(context));
    return SCPI_RES_ERR;
}

// new scpi object
scpi_object_s* newScpi(void){
	scpi_object_s* obj = (scpi_object_s*)malloc(sizeof(scpi_object_s));
	memset(obj, 0, sizeof(scpi_object_s));

	obj->interf.error = scpi_error_callback;
	obj->interf.write = scpi_write;
	obj->interf.control = scpi_control;
	obj->interf.flush = scpi_flush;
	obj->interf.reset = scpi_reset;

	SCPI_Init(&obj->ctx,
		obj->cmds,
		&obj->interf,
		scpi_units_def,
		0, 0, 0, 0, 0,
		0,
		obj->errq, sizeof(obj->errq)/sizeof(scpi_error_t));
	return obj;
}

// add scpi command
scpi_result_t addScpi(scpi_object_s* s, _GoString_ pattern, uint64_t uid){
	scpi_command_t* cmd = &s->cmds[s->cmd_idx];

	// copy pattern
	cmd->pattern = (char*)malloc(_GoStringLen(pattern) + 1);
	memcpy(cmd->pattern, _GoStringPtr(pattern), _GoStringLen(pattern));
	*(cmd->pattern + _GoStringLen(pattern))  = '\0';

	// callback function
	cmd->callback = hdl_scpi;

	// uid
	cmd->tag = uid;

	// next
	s->cmd_idx += 1;
	return SCPI_RES_OK;
}

void cleanScpi(scpi_object_s* s){
	// free all command pattern
	// xxx

	// free
	free(s);
}

int parseSCPI(scpi_object_s* s, _GoString_ cmd){
	return SCPI_Parse(&s->ctx, (char*)_GoStringPtr(cmd), _GoStringLen(cmd));
}
*/
import "C"

type Scpi struct {
	o *C.struct_scpi_object_s
}

func NewScpi() *Scpi {
	return &Scpi{C.newScpi()}
}

func (s *Scpi) AddScpi(pattern string) {
	C.addScpi(s.o, pattern, 123)
}

func (s *Scpi) ParseScpi(cmd string) {
	C.parseSCPI(s.o, cmd)
}

func (s *Scpi) Clean() {
	C.cleanScpi(s.o)
}

func main() {
	scpi := NewScpi()

	scpi.AddScpi("*RST")
	scpi.ParseScpi("*RST")

	scpi.Clean()
}
