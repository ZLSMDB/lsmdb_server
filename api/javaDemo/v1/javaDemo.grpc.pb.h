// Generated by the gRPC C++ plugin.
// If you make any local change, they will be lost.
// source: javaDemo/v1/javaDemo.proto
#ifndef GRPC_javaDemo_2fv1_2fjavaDemo_2eproto__INCLUDED
#define GRPC_javaDemo_2fv1_2fjavaDemo_2eproto__INCLUDED

#include "javaDemo/v1/javaDemo.pb.h"

#include <functional>
#include <grpcpp/generic/async_generic_service.h>
#include <grpcpp/support/async_stream.h>
#include <grpcpp/support/async_unary_call.h>
#include <grpcpp/support/client_callback.h>
#include <grpcpp/client_context.h>
#include <grpcpp/completion_queue.h>
#include <grpcpp/support/message_allocator.h>
#include <grpcpp/support/method_handler.h>
#include <grpcpp/impl/proto_utils.h>
#include <grpcpp/impl/rpc_method.h>
#include <grpcpp/support/server_callback.h>
#include <grpcpp/impl/server_callback_handlers.h>
#include <grpcpp/server_context.h>
#include <grpcpp/impl/service_type.h>
#include <grpcpp/support/status.h>
#include <grpcpp/support/stub_options.h>
#include <grpcpp/support/sync_stream.h>

namespace javaDemo {
namespace v1 {

// 定义接口
class javaDemoService final {
 public:
  static constexpr char const* service_full_name() {
    return "javaDemo.v1.javaDemoService";
  }
  class StubInterface {
   public:
    virtual ~StubInterface() {}
    // 一个简单的rpc
    virtual ::grpc::Status HelloWorld(::grpc::ClientContext* context, const ::javaDemo::v1::HelloRequest& request, ::javaDemo::v1::HelloResponse* response) = 0;
    std::unique_ptr< ::grpc::ClientAsyncResponseReaderInterface< ::javaDemo::v1::HelloResponse>> AsyncHelloWorld(::grpc::ClientContext* context, const ::javaDemo::v1::HelloRequest& request, ::grpc::CompletionQueue* cq) {
      return std::unique_ptr< ::grpc::ClientAsyncResponseReaderInterface< ::javaDemo::v1::HelloResponse>>(AsyncHelloWorldRaw(context, request, cq));
    }
    std::unique_ptr< ::grpc::ClientAsyncResponseReaderInterface< ::javaDemo::v1::HelloResponse>> PrepareAsyncHelloWorld(::grpc::ClientContext* context, const ::javaDemo::v1::HelloRequest& request, ::grpc::CompletionQueue* cq) {
      return std::unique_ptr< ::grpc::ClientAsyncResponseReaderInterface< ::javaDemo::v1::HelloResponse>>(PrepareAsyncHelloWorldRaw(context, request, cq));
    }
    // 一个客户端流式rpc
    std::unique_ptr< ::grpc::ClientWriterInterface< ::javaDemo::v1::HelloRequest>> HelloWorldClientStream(::grpc::ClientContext* context, ::javaDemo::v1::HelloResponse* response) {
      return std::unique_ptr< ::grpc::ClientWriterInterface< ::javaDemo::v1::HelloRequest>>(HelloWorldClientStreamRaw(context, response));
    }
    std::unique_ptr< ::grpc::ClientAsyncWriterInterface< ::javaDemo::v1::HelloRequest>> AsyncHelloWorldClientStream(::grpc::ClientContext* context, ::javaDemo::v1::HelloResponse* response, ::grpc::CompletionQueue* cq, void* tag) {
      return std::unique_ptr< ::grpc::ClientAsyncWriterInterface< ::javaDemo::v1::HelloRequest>>(AsyncHelloWorldClientStreamRaw(context, response, cq, tag));
    }
    std::unique_ptr< ::grpc::ClientAsyncWriterInterface< ::javaDemo::v1::HelloRequest>> PrepareAsyncHelloWorldClientStream(::grpc::ClientContext* context, ::javaDemo::v1::HelloResponse* response, ::grpc::CompletionQueue* cq) {
      return std::unique_ptr< ::grpc::ClientAsyncWriterInterface< ::javaDemo::v1::HelloRequest>>(PrepareAsyncHelloWorldClientStreamRaw(context, response, cq));
    }
    // 一个客户端和服务器端双向流式rpc
    std::unique_ptr< ::grpc::ClientReaderWriterInterface< ::javaDemo::v1::HelloRequest, ::javaDemo::v1::HelloResponse>> HelloWorldClientAndServerStream(::grpc::ClientContext* context) {
      return std::unique_ptr< ::grpc::ClientReaderWriterInterface< ::javaDemo::v1::HelloRequest, ::javaDemo::v1::HelloResponse>>(HelloWorldClientAndServerStreamRaw(context));
    }
    std::unique_ptr< ::grpc::ClientAsyncReaderWriterInterface< ::javaDemo::v1::HelloRequest, ::javaDemo::v1::HelloResponse>> AsyncHelloWorldClientAndServerStream(::grpc::ClientContext* context, ::grpc::CompletionQueue* cq, void* tag) {
      return std::unique_ptr< ::grpc::ClientAsyncReaderWriterInterface< ::javaDemo::v1::HelloRequest, ::javaDemo::v1::HelloResponse>>(AsyncHelloWorldClientAndServerStreamRaw(context, cq, tag));
    }
    std::unique_ptr< ::grpc::ClientAsyncReaderWriterInterface< ::javaDemo::v1::HelloRequest, ::javaDemo::v1::HelloResponse>> PrepareAsyncHelloWorldClientAndServerStream(::grpc::ClientContext* context, ::grpc::CompletionQueue* cq) {
      return std::unique_ptr< ::grpc::ClientAsyncReaderWriterInterface< ::javaDemo::v1::HelloRequest, ::javaDemo::v1::HelloResponse>>(PrepareAsyncHelloWorldClientAndServerStreamRaw(context, cq));
    }
    class async_interface {
     public:
      virtual ~async_interface() {}
      // 一个简单的rpc
      virtual void HelloWorld(::grpc::ClientContext* context, const ::javaDemo::v1::HelloRequest* request, ::javaDemo::v1::HelloResponse* response, std::function<void(::grpc::Status)>) = 0;
      virtual void HelloWorld(::grpc::ClientContext* context, const ::javaDemo::v1::HelloRequest* request, ::javaDemo::v1::HelloResponse* response, ::grpc::ClientUnaryReactor* reactor) = 0;
      // 一个客户端流式rpc
      virtual void HelloWorldClientStream(::grpc::ClientContext* context, ::javaDemo::v1::HelloResponse* response, ::grpc::ClientWriteReactor< ::javaDemo::v1::HelloRequest>* reactor) = 0;
      // 一个客户端和服务器端双向流式rpc
      virtual void HelloWorldClientAndServerStream(::grpc::ClientContext* context, ::grpc::ClientBidiReactor< ::javaDemo::v1::HelloRequest,::javaDemo::v1::HelloResponse>* reactor) = 0;
    };
    typedef class async_interface experimental_async_interface;
    virtual class async_interface* async() { return nullptr; }
    class async_interface* experimental_async() { return async(); }
   private:
    virtual ::grpc::ClientAsyncResponseReaderInterface< ::javaDemo::v1::HelloResponse>* AsyncHelloWorldRaw(::grpc::ClientContext* context, const ::javaDemo::v1::HelloRequest& request, ::grpc::CompletionQueue* cq) = 0;
    virtual ::grpc::ClientAsyncResponseReaderInterface< ::javaDemo::v1::HelloResponse>* PrepareAsyncHelloWorldRaw(::grpc::ClientContext* context, const ::javaDemo::v1::HelloRequest& request, ::grpc::CompletionQueue* cq) = 0;
    virtual ::grpc::ClientWriterInterface< ::javaDemo::v1::HelloRequest>* HelloWorldClientStreamRaw(::grpc::ClientContext* context, ::javaDemo::v1::HelloResponse* response) = 0;
    virtual ::grpc::ClientAsyncWriterInterface< ::javaDemo::v1::HelloRequest>* AsyncHelloWorldClientStreamRaw(::grpc::ClientContext* context, ::javaDemo::v1::HelloResponse* response, ::grpc::CompletionQueue* cq, void* tag) = 0;
    virtual ::grpc::ClientAsyncWriterInterface< ::javaDemo::v1::HelloRequest>* PrepareAsyncHelloWorldClientStreamRaw(::grpc::ClientContext* context, ::javaDemo::v1::HelloResponse* response, ::grpc::CompletionQueue* cq) = 0;
    virtual ::grpc::ClientReaderWriterInterface< ::javaDemo::v1::HelloRequest, ::javaDemo::v1::HelloResponse>* HelloWorldClientAndServerStreamRaw(::grpc::ClientContext* context) = 0;
    virtual ::grpc::ClientAsyncReaderWriterInterface< ::javaDemo::v1::HelloRequest, ::javaDemo::v1::HelloResponse>* AsyncHelloWorldClientAndServerStreamRaw(::grpc::ClientContext* context, ::grpc::CompletionQueue* cq, void* tag) = 0;
    virtual ::grpc::ClientAsyncReaderWriterInterface< ::javaDemo::v1::HelloRequest, ::javaDemo::v1::HelloResponse>* PrepareAsyncHelloWorldClientAndServerStreamRaw(::grpc::ClientContext* context, ::grpc::CompletionQueue* cq) = 0;
  };
  class Stub final : public StubInterface {
   public:
    Stub(const std::shared_ptr< ::grpc::ChannelInterface>& channel, const ::grpc::StubOptions& options = ::grpc::StubOptions());
    ::grpc::Status HelloWorld(::grpc::ClientContext* context, const ::javaDemo::v1::HelloRequest& request, ::javaDemo::v1::HelloResponse* response) override;
    std::unique_ptr< ::grpc::ClientAsyncResponseReader< ::javaDemo::v1::HelloResponse>> AsyncHelloWorld(::grpc::ClientContext* context, const ::javaDemo::v1::HelloRequest& request, ::grpc::CompletionQueue* cq) {
      return std::unique_ptr< ::grpc::ClientAsyncResponseReader< ::javaDemo::v1::HelloResponse>>(AsyncHelloWorldRaw(context, request, cq));
    }
    std::unique_ptr< ::grpc::ClientAsyncResponseReader< ::javaDemo::v1::HelloResponse>> PrepareAsyncHelloWorld(::grpc::ClientContext* context, const ::javaDemo::v1::HelloRequest& request, ::grpc::CompletionQueue* cq) {
      return std::unique_ptr< ::grpc::ClientAsyncResponseReader< ::javaDemo::v1::HelloResponse>>(PrepareAsyncHelloWorldRaw(context, request, cq));
    }
    std::unique_ptr< ::grpc::ClientWriter< ::javaDemo::v1::HelloRequest>> HelloWorldClientStream(::grpc::ClientContext* context, ::javaDemo::v1::HelloResponse* response) {
      return std::unique_ptr< ::grpc::ClientWriter< ::javaDemo::v1::HelloRequest>>(HelloWorldClientStreamRaw(context, response));
    }
    std::unique_ptr< ::grpc::ClientAsyncWriter< ::javaDemo::v1::HelloRequest>> AsyncHelloWorldClientStream(::grpc::ClientContext* context, ::javaDemo::v1::HelloResponse* response, ::grpc::CompletionQueue* cq, void* tag) {
      return std::unique_ptr< ::grpc::ClientAsyncWriter< ::javaDemo::v1::HelloRequest>>(AsyncHelloWorldClientStreamRaw(context, response, cq, tag));
    }
    std::unique_ptr< ::grpc::ClientAsyncWriter< ::javaDemo::v1::HelloRequest>> PrepareAsyncHelloWorldClientStream(::grpc::ClientContext* context, ::javaDemo::v1::HelloResponse* response, ::grpc::CompletionQueue* cq) {
      return std::unique_ptr< ::grpc::ClientAsyncWriter< ::javaDemo::v1::HelloRequest>>(PrepareAsyncHelloWorldClientStreamRaw(context, response, cq));
    }
    std::unique_ptr< ::grpc::ClientReaderWriter< ::javaDemo::v1::HelloRequest, ::javaDemo::v1::HelloResponse>> HelloWorldClientAndServerStream(::grpc::ClientContext* context) {
      return std::unique_ptr< ::grpc::ClientReaderWriter< ::javaDemo::v1::HelloRequest, ::javaDemo::v1::HelloResponse>>(HelloWorldClientAndServerStreamRaw(context));
    }
    std::unique_ptr<  ::grpc::ClientAsyncReaderWriter< ::javaDemo::v1::HelloRequest, ::javaDemo::v1::HelloResponse>> AsyncHelloWorldClientAndServerStream(::grpc::ClientContext* context, ::grpc::CompletionQueue* cq, void* tag) {
      return std::unique_ptr< ::grpc::ClientAsyncReaderWriter< ::javaDemo::v1::HelloRequest, ::javaDemo::v1::HelloResponse>>(AsyncHelloWorldClientAndServerStreamRaw(context, cq, tag));
    }
    std::unique_ptr<  ::grpc::ClientAsyncReaderWriter< ::javaDemo::v1::HelloRequest, ::javaDemo::v1::HelloResponse>> PrepareAsyncHelloWorldClientAndServerStream(::grpc::ClientContext* context, ::grpc::CompletionQueue* cq) {
      return std::unique_ptr< ::grpc::ClientAsyncReaderWriter< ::javaDemo::v1::HelloRequest, ::javaDemo::v1::HelloResponse>>(PrepareAsyncHelloWorldClientAndServerStreamRaw(context, cq));
    }
    class async final :
      public StubInterface::async_interface {
     public:
      void HelloWorld(::grpc::ClientContext* context, const ::javaDemo::v1::HelloRequest* request, ::javaDemo::v1::HelloResponse* response, std::function<void(::grpc::Status)>) override;
      void HelloWorld(::grpc::ClientContext* context, const ::javaDemo::v1::HelloRequest* request, ::javaDemo::v1::HelloResponse* response, ::grpc::ClientUnaryReactor* reactor) override;
      void HelloWorldClientStream(::grpc::ClientContext* context, ::javaDemo::v1::HelloResponse* response, ::grpc::ClientWriteReactor< ::javaDemo::v1::HelloRequest>* reactor) override;
      void HelloWorldClientAndServerStream(::grpc::ClientContext* context, ::grpc::ClientBidiReactor< ::javaDemo::v1::HelloRequest,::javaDemo::v1::HelloResponse>* reactor) override;
     private:
      friend class Stub;
      explicit async(Stub* stub): stub_(stub) { }
      Stub* stub() { return stub_; }
      Stub* stub_;
    };
    class async* async() override { return &async_stub_; }

   private:
    std::shared_ptr< ::grpc::ChannelInterface> channel_;
    class async async_stub_{this};
    ::grpc::ClientAsyncResponseReader< ::javaDemo::v1::HelloResponse>* AsyncHelloWorldRaw(::grpc::ClientContext* context, const ::javaDemo::v1::HelloRequest& request, ::grpc::CompletionQueue* cq) override;
    ::grpc::ClientAsyncResponseReader< ::javaDemo::v1::HelloResponse>* PrepareAsyncHelloWorldRaw(::grpc::ClientContext* context, const ::javaDemo::v1::HelloRequest& request, ::grpc::CompletionQueue* cq) override;
    ::grpc::ClientWriter< ::javaDemo::v1::HelloRequest>* HelloWorldClientStreamRaw(::grpc::ClientContext* context, ::javaDemo::v1::HelloResponse* response) override;
    ::grpc::ClientAsyncWriter< ::javaDemo::v1::HelloRequest>* AsyncHelloWorldClientStreamRaw(::grpc::ClientContext* context, ::javaDemo::v1::HelloResponse* response, ::grpc::CompletionQueue* cq, void* tag) override;
    ::grpc::ClientAsyncWriter< ::javaDemo::v1::HelloRequest>* PrepareAsyncHelloWorldClientStreamRaw(::grpc::ClientContext* context, ::javaDemo::v1::HelloResponse* response, ::grpc::CompletionQueue* cq) override;
    ::grpc::ClientReaderWriter< ::javaDemo::v1::HelloRequest, ::javaDemo::v1::HelloResponse>* HelloWorldClientAndServerStreamRaw(::grpc::ClientContext* context) override;
    ::grpc::ClientAsyncReaderWriter< ::javaDemo::v1::HelloRequest, ::javaDemo::v1::HelloResponse>* AsyncHelloWorldClientAndServerStreamRaw(::grpc::ClientContext* context, ::grpc::CompletionQueue* cq, void* tag) override;
    ::grpc::ClientAsyncReaderWriter< ::javaDemo::v1::HelloRequest, ::javaDemo::v1::HelloResponse>* PrepareAsyncHelloWorldClientAndServerStreamRaw(::grpc::ClientContext* context, ::grpc::CompletionQueue* cq) override;
    const ::grpc::internal::RpcMethod rpcmethod_HelloWorld_;
    const ::grpc::internal::RpcMethod rpcmethod_HelloWorldClientStream_;
    const ::grpc::internal::RpcMethod rpcmethod_HelloWorldClientAndServerStream_;
  };
  static std::unique_ptr<Stub> NewStub(const std::shared_ptr< ::grpc::ChannelInterface>& channel, const ::grpc::StubOptions& options = ::grpc::StubOptions());

  class Service : public ::grpc::Service {
   public:
    Service();
    virtual ~Service();
    // 一个简单的rpc
    virtual ::grpc::Status HelloWorld(::grpc::ServerContext* context, const ::javaDemo::v1::HelloRequest* request, ::javaDemo::v1::HelloResponse* response);
    // 一个客户端流式rpc
    virtual ::grpc::Status HelloWorldClientStream(::grpc::ServerContext* context, ::grpc::ServerReader< ::javaDemo::v1::HelloRequest>* reader, ::javaDemo::v1::HelloResponse* response);
    // 一个客户端和服务器端双向流式rpc
    virtual ::grpc::Status HelloWorldClientAndServerStream(::grpc::ServerContext* context, ::grpc::ServerReaderWriter< ::javaDemo::v1::HelloResponse, ::javaDemo::v1::HelloRequest>* stream);
  };
  template <class BaseClass>
  class WithAsyncMethod_HelloWorld : public BaseClass {
   private:
    void BaseClassMustBeDerivedFromService(const Service* /*service*/) {}
   public:
    WithAsyncMethod_HelloWorld() {
      ::grpc::Service::MarkMethodAsync(0);
    }
    ~WithAsyncMethod_HelloWorld() override {
      BaseClassMustBeDerivedFromService(this);
    }
    // disable synchronous version of this method
    ::grpc::Status HelloWorld(::grpc::ServerContext* /*context*/, const ::javaDemo::v1::HelloRequest* /*request*/, ::javaDemo::v1::HelloResponse* /*response*/) override {
      abort();
      return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
    }
    void RequestHelloWorld(::grpc::ServerContext* context, ::javaDemo::v1::HelloRequest* request, ::grpc::ServerAsyncResponseWriter< ::javaDemo::v1::HelloResponse>* response, ::grpc::CompletionQueue* new_call_cq, ::grpc::ServerCompletionQueue* notification_cq, void *tag) {
      ::grpc::Service::RequestAsyncUnary(0, context, request, response, new_call_cq, notification_cq, tag);
    }
  };
  template <class BaseClass>
  class WithAsyncMethod_HelloWorldClientStream : public BaseClass {
   private:
    void BaseClassMustBeDerivedFromService(const Service* /*service*/) {}
   public:
    WithAsyncMethod_HelloWorldClientStream() {
      ::grpc::Service::MarkMethodAsync(1);
    }
    ~WithAsyncMethod_HelloWorldClientStream() override {
      BaseClassMustBeDerivedFromService(this);
    }
    // disable synchronous version of this method
    ::grpc::Status HelloWorldClientStream(::grpc::ServerContext* /*context*/, ::grpc::ServerReader< ::javaDemo::v1::HelloRequest>* /*reader*/, ::javaDemo::v1::HelloResponse* /*response*/) override {
      abort();
      return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
    }
    void RequestHelloWorldClientStream(::grpc::ServerContext* context, ::grpc::ServerAsyncReader< ::javaDemo::v1::HelloResponse, ::javaDemo::v1::HelloRequest>* reader, ::grpc::CompletionQueue* new_call_cq, ::grpc::ServerCompletionQueue* notification_cq, void *tag) {
      ::grpc::Service::RequestAsyncClientStreaming(1, context, reader, new_call_cq, notification_cq, tag);
    }
  };
  template <class BaseClass>
  class WithAsyncMethod_HelloWorldClientAndServerStream : public BaseClass {
   private:
    void BaseClassMustBeDerivedFromService(const Service* /*service*/) {}
   public:
    WithAsyncMethod_HelloWorldClientAndServerStream() {
      ::grpc::Service::MarkMethodAsync(2);
    }
    ~WithAsyncMethod_HelloWorldClientAndServerStream() override {
      BaseClassMustBeDerivedFromService(this);
    }
    // disable synchronous version of this method
    ::grpc::Status HelloWorldClientAndServerStream(::grpc::ServerContext* /*context*/, ::grpc::ServerReaderWriter< ::javaDemo::v1::HelloResponse, ::javaDemo::v1::HelloRequest>* /*stream*/)  override {
      abort();
      return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
    }
    void RequestHelloWorldClientAndServerStream(::grpc::ServerContext* context, ::grpc::ServerAsyncReaderWriter< ::javaDemo::v1::HelloResponse, ::javaDemo::v1::HelloRequest>* stream, ::grpc::CompletionQueue* new_call_cq, ::grpc::ServerCompletionQueue* notification_cq, void *tag) {
      ::grpc::Service::RequestAsyncBidiStreaming(2, context, stream, new_call_cq, notification_cq, tag);
    }
  };
  typedef WithAsyncMethod_HelloWorld<WithAsyncMethod_HelloWorldClientStream<WithAsyncMethod_HelloWorldClientAndServerStream<Service > > > AsyncService;
  template <class BaseClass>
  class WithCallbackMethod_HelloWorld : public BaseClass {
   private:
    void BaseClassMustBeDerivedFromService(const Service* /*service*/) {}
   public:
    WithCallbackMethod_HelloWorld() {
      ::grpc::Service::MarkMethodCallback(0,
          new ::grpc::internal::CallbackUnaryHandler< ::javaDemo::v1::HelloRequest, ::javaDemo::v1::HelloResponse>(
            [this](
                   ::grpc::CallbackServerContext* context, const ::javaDemo::v1::HelloRequest* request, ::javaDemo::v1::HelloResponse* response) { return this->HelloWorld(context, request, response); }));}
    void SetMessageAllocatorFor_HelloWorld(
        ::grpc::MessageAllocator< ::javaDemo::v1::HelloRequest, ::javaDemo::v1::HelloResponse>* allocator) {
      ::grpc::internal::MethodHandler* const handler = ::grpc::Service::GetHandler(0);
      static_cast<::grpc::internal::CallbackUnaryHandler< ::javaDemo::v1::HelloRequest, ::javaDemo::v1::HelloResponse>*>(handler)
              ->SetMessageAllocator(allocator);
    }
    ~WithCallbackMethod_HelloWorld() override {
      BaseClassMustBeDerivedFromService(this);
    }
    // disable synchronous version of this method
    ::grpc::Status HelloWorld(::grpc::ServerContext* /*context*/, const ::javaDemo::v1::HelloRequest* /*request*/, ::javaDemo::v1::HelloResponse* /*response*/) override {
      abort();
      return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
    }
    virtual ::grpc::ServerUnaryReactor* HelloWorld(
      ::grpc::CallbackServerContext* /*context*/, const ::javaDemo::v1::HelloRequest* /*request*/, ::javaDemo::v1::HelloResponse* /*response*/)  { return nullptr; }
  };
  template <class BaseClass>
  class WithCallbackMethod_HelloWorldClientStream : public BaseClass {
   private:
    void BaseClassMustBeDerivedFromService(const Service* /*service*/) {}
   public:
    WithCallbackMethod_HelloWorldClientStream() {
      ::grpc::Service::MarkMethodCallback(1,
          new ::grpc::internal::CallbackClientStreamingHandler< ::javaDemo::v1::HelloRequest, ::javaDemo::v1::HelloResponse>(
            [this](
                   ::grpc::CallbackServerContext* context, ::javaDemo::v1::HelloResponse* response) { return this->HelloWorldClientStream(context, response); }));
    }
    ~WithCallbackMethod_HelloWorldClientStream() override {
      BaseClassMustBeDerivedFromService(this);
    }
    // disable synchronous version of this method
    ::grpc::Status HelloWorldClientStream(::grpc::ServerContext* /*context*/, ::grpc::ServerReader< ::javaDemo::v1::HelloRequest>* /*reader*/, ::javaDemo::v1::HelloResponse* /*response*/) override {
      abort();
      return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
    }
    virtual ::grpc::ServerReadReactor< ::javaDemo::v1::HelloRequest>* HelloWorldClientStream(
      ::grpc::CallbackServerContext* /*context*/, ::javaDemo::v1::HelloResponse* /*response*/)  { return nullptr; }
  };
  template <class BaseClass>
  class WithCallbackMethod_HelloWorldClientAndServerStream : public BaseClass {
   private:
    void BaseClassMustBeDerivedFromService(const Service* /*service*/) {}
   public:
    WithCallbackMethod_HelloWorldClientAndServerStream() {
      ::grpc::Service::MarkMethodCallback(2,
          new ::grpc::internal::CallbackBidiHandler< ::javaDemo::v1::HelloRequest, ::javaDemo::v1::HelloResponse>(
            [this](
                   ::grpc::CallbackServerContext* context) { return this->HelloWorldClientAndServerStream(context); }));
    }
    ~WithCallbackMethod_HelloWorldClientAndServerStream() override {
      BaseClassMustBeDerivedFromService(this);
    }
    // disable synchronous version of this method
    ::grpc::Status HelloWorldClientAndServerStream(::grpc::ServerContext* /*context*/, ::grpc::ServerReaderWriter< ::javaDemo::v1::HelloResponse, ::javaDemo::v1::HelloRequest>* /*stream*/)  override {
      abort();
      return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
    }
    virtual ::grpc::ServerBidiReactor< ::javaDemo::v1::HelloRequest, ::javaDemo::v1::HelloResponse>* HelloWorldClientAndServerStream(
      ::grpc::CallbackServerContext* /*context*/)
      { return nullptr; }
  };
  typedef WithCallbackMethod_HelloWorld<WithCallbackMethod_HelloWorldClientStream<WithCallbackMethod_HelloWorldClientAndServerStream<Service > > > CallbackService;
  typedef CallbackService ExperimentalCallbackService;
  template <class BaseClass>
  class WithGenericMethod_HelloWorld : public BaseClass {
   private:
    void BaseClassMustBeDerivedFromService(const Service* /*service*/) {}
   public:
    WithGenericMethod_HelloWorld() {
      ::grpc::Service::MarkMethodGeneric(0);
    }
    ~WithGenericMethod_HelloWorld() override {
      BaseClassMustBeDerivedFromService(this);
    }
    // disable synchronous version of this method
    ::grpc::Status HelloWorld(::grpc::ServerContext* /*context*/, const ::javaDemo::v1::HelloRequest* /*request*/, ::javaDemo::v1::HelloResponse* /*response*/) override {
      abort();
      return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
    }
  };
  template <class BaseClass>
  class WithGenericMethod_HelloWorldClientStream : public BaseClass {
   private:
    void BaseClassMustBeDerivedFromService(const Service* /*service*/) {}
   public:
    WithGenericMethod_HelloWorldClientStream() {
      ::grpc::Service::MarkMethodGeneric(1);
    }
    ~WithGenericMethod_HelloWorldClientStream() override {
      BaseClassMustBeDerivedFromService(this);
    }
    // disable synchronous version of this method
    ::grpc::Status HelloWorldClientStream(::grpc::ServerContext* /*context*/, ::grpc::ServerReader< ::javaDemo::v1::HelloRequest>* /*reader*/, ::javaDemo::v1::HelloResponse* /*response*/) override {
      abort();
      return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
    }
  };
  template <class BaseClass>
  class WithGenericMethod_HelloWorldClientAndServerStream : public BaseClass {
   private:
    void BaseClassMustBeDerivedFromService(const Service* /*service*/) {}
   public:
    WithGenericMethod_HelloWorldClientAndServerStream() {
      ::grpc::Service::MarkMethodGeneric(2);
    }
    ~WithGenericMethod_HelloWorldClientAndServerStream() override {
      BaseClassMustBeDerivedFromService(this);
    }
    // disable synchronous version of this method
    ::grpc::Status HelloWorldClientAndServerStream(::grpc::ServerContext* /*context*/, ::grpc::ServerReaderWriter< ::javaDemo::v1::HelloResponse, ::javaDemo::v1::HelloRequest>* /*stream*/)  override {
      abort();
      return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
    }
  };
  template <class BaseClass>
  class WithRawMethod_HelloWorld : public BaseClass {
   private:
    void BaseClassMustBeDerivedFromService(const Service* /*service*/) {}
   public:
    WithRawMethod_HelloWorld() {
      ::grpc::Service::MarkMethodRaw(0);
    }
    ~WithRawMethod_HelloWorld() override {
      BaseClassMustBeDerivedFromService(this);
    }
    // disable synchronous version of this method
    ::grpc::Status HelloWorld(::grpc::ServerContext* /*context*/, const ::javaDemo::v1::HelloRequest* /*request*/, ::javaDemo::v1::HelloResponse* /*response*/) override {
      abort();
      return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
    }
    void RequestHelloWorld(::grpc::ServerContext* context, ::grpc::ByteBuffer* request, ::grpc::ServerAsyncResponseWriter< ::grpc::ByteBuffer>* response, ::grpc::CompletionQueue* new_call_cq, ::grpc::ServerCompletionQueue* notification_cq, void *tag) {
      ::grpc::Service::RequestAsyncUnary(0, context, request, response, new_call_cq, notification_cq, tag);
    }
  };
  template <class BaseClass>
  class WithRawMethod_HelloWorldClientStream : public BaseClass {
   private:
    void BaseClassMustBeDerivedFromService(const Service* /*service*/) {}
   public:
    WithRawMethod_HelloWorldClientStream() {
      ::grpc::Service::MarkMethodRaw(1);
    }
    ~WithRawMethod_HelloWorldClientStream() override {
      BaseClassMustBeDerivedFromService(this);
    }
    // disable synchronous version of this method
    ::grpc::Status HelloWorldClientStream(::grpc::ServerContext* /*context*/, ::grpc::ServerReader< ::javaDemo::v1::HelloRequest>* /*reader*/, ::javaDemo::v1::HelloResponse* /*response*/) override {
      abort();
      return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
    }
    void RequestHelloWorldClientStream(::grpc::ServerContext* context, ::grpc::ServerAsyncReader< ::grpc::ByteBuffer, ::grpc::ByteBuffer>* reader, ::grpc::CompletionQueue* new_call_cq, ::grpc::ServerCompletionQueue* notification_cq, void *tag) {
      ::grpc::Service::RequestAsyncClientStreaming(1, context, reader, new_call_cq, notification_cq, tag);
    }
  };
  template <class BaseClass>
  class WithRawMethod_HelloWorldClientAndServerStream : public BaseClass {
   private:
    void BaseClassMustBeDerivedFromService(const Service* /*service*/) {}
   public:
    WithRawMethod_HelloWorldClientAndServerStream() {
      ::grpc::Service::MarkMethodRaw(2);
    }
    ~WithRawMethod_HelloWorldClientAndServerStream() override {
      BaseClassMustBeDerivedFromService(this);
    }
    // disable synchronous version of this method
    ::grpc::Status HelloWorldClientAndServerStream(::grpc::ServerContext* /*context*/, ::grpc::ServerReaderWriter< ::javaDemo::v1::HelloResponse, ::javaDemo::v1::HelloRequest>* /*stream*/)  override {
      abort();
      return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
    }
    void RequestHelloWorldClientAndServerStream(::grpc::ServerContext* context, ::grpc::ServerAsyncReaderWriter< ::grpc::ByteBuffer, ::grpc::ByteBuffer>* stream, ::grpc::CompletionQueue* new_call_cq, ::grpc::ServerCompletionQueue* notification_cq, void *tag) {
      ::grpc::Service::RequestAsyncBidiStreaming(2, context, stream, new_call_cq, notification_cq, tag);
    }
  };
  template <class BaseClass>
  class WithRawCallbackMethod_HelloWorld : public BaseClass {
   private:
    void BaseClassMustBeDerivedFromService(const Service* /*service*/) {}
   public:
    WithRawCallbackMethod_HelloWorld() {
      ::grpc::Service::MarkMethodRawCallback(0,
          new ::grpc::internal::CallbackUnaryHandler< ::grpc::ByteBuffer, ::grpc::ByteBuffer>(
            [this](
                   ::grpc::CallbackServerContext* context, const ::grpc::ByteBuffer* request, ::grpc::ByteBuffer* response) { return this->HelloWorld(context, request, response); }));
    }
    ~WithRawCallbackMethod_HelloWorld() override {
      BaseClassMustBeDerivedFromService(this);
    }
    // disable synchronous version of this method
    ::grpc::Status HelloWorld(::grpc::ServerContext* /*context*/, const ::javaDemo::v1::HelloRequest* /*request*/, ::javaDemo::v1::HelloResponse* /*response*/) override {
      abort();
      return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
    }
    virtual ::grpc::ServerUnaryReactor* HelloWorld(
      ::grpc::CallbackServerContext* /*context*/, const ::grpc::ByteBuffer* /*request*/, ::grpc::ByteBuffer* /*response*/)  { return nullptr; }
  };
  template <class BaseClass>
  class WithRawCallbackMethod_HelloWorldClientStream : public BaseClass {
   private:
    void BaseClassMustBeDerivedFromService(const Service* /*service*/) {}
   public:
    WithRawCallbackMethod_HelloWorldClientStream() {
      ::grpc::Service::MarkMethodRawCallback(1,
          new ::grpc::internal::CallbackClientStreamingHandler< ::grpc::ByteBuffer, ::grpc::ByteBuffer>(
            [this](
                   ::grpc::CallbackServerContext* context, ::grpc::ByteBuffer* response) { return this->HelloWorldClientStream(context, response); }));
    }
    ~WithRawCallbackMethod_HelloWorldClientStream() override {
      BaseClassMustBeDerivedFromService(this);
    }
    // disable synchronous version of this method
    ::grpc::Status HelloWorldClientStream(::grpc::ServerContext* /*context*/, ::grpc::ServerReader< ::javaDemo::v1::HelloRequest>* /*reader*/, ::javaDemo::v1::HelloResponse* /*response*/) override {
      abort();
      return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
    }
    virtual ::grpc::ServerReadReactor< ::grpc::ByteBuffer>* HelloWorldClientStream(
      ::grpc::CallbackServerContext* /*context*/, ::grpc::ByteBuffer* /*response*/)  { return nullptr; }
  };
  template <class BaseClass>
  class WithRawCallbackMethod_HelloWorldClientAndServerStream : public BaseClass {
   private:
    void BaseClassMustBeDerivedFromService(const Service* /*service*/) {}
   public:
    WithRawCallbackMethod_HelloWorldClientAndServerStream() {
      ::grpc::Service::MarkMethodRawCallback(2,
          new ::grpc::internal::CallbackBidiHandler< ::grpc::ByteBuffer, ::grpc::ByteBuffer>(
            [this](
                   ::grpc::CallbackServerContext* context) { return this->HelloWorldClientAndServerStream(context); }));
    }
    ~WithRawCallbackMethod_HelloWorldClientAndServerStream() override {
      BaseClassMustBeDerivedFromService(this);
    }
    // disable synchronous version of this method
    ::grpc::Status HelloWorldClientAndServerStream(::grpc::ServerContext* /*context*/, ::grpc::ServerReaderWriter< ::javaDemo::v1::HelloResponse, ::javaDemo::v1::HelloRequest>* /*stream*/)  override {
      abort();
      return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
    }
    virtual ::grpc::ServerBidiReactor< ::grpc::ByteBuffer, ::grpc::ByteBuffer>* HelloWorldClientAndServerStream(
      ::grpc::CallbackServerContext* /*context*/)
      { return nullptr; }
  };
  template <class BaseClass>
  class WithStreamedUnaryMethod_HelloWorld : public BaseClass {
   private:
    void BaseClassMustBeDerivedFromService(const Service* /*service*/) {}
   public:
    WithStreamedUnaryMethod_HelloWorld() {
      ::grpc::Service::MarkMethodStreamed(0,
        new ::grpc::internal::StreamedUnaryHandler<
          ::javaDemo::v1::HelloRequest, ::javaDemo::v1::HelloResponse>(
            [this](::grpc::ServerContext* context,
                   ::grpc::ServerUnaryStreamer<
                     ::javaDemo::v1::HelloRequest, ::javaDemo::v1::HelloResponse>* streamer) {
                       return this->StreamedHelloWorld(context,
                         streamer);
                  }));
    }
    ~WithStreamedUnaryMethod_HelloWorld() override {
      BaseClassMustBeDerivedFromService(this);
    }
    // disable regular version of this method
    ::grpc::Status HelloWorld(::grpc::ServerContext* /*context*/, const ::javaDemo::v1::HelloRequest* /*request*/, ::javaDemo::v1::HelloResponse* /*response*/) override {
      abort();
      return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
    }
    // replace default version of method with streamed unary
    virtual ::grpc::Status StreamedHelloWorld(::grpc::ServerContext* context, ::grpc::ServerUnaryStreamer< ::javaDemo::v1::HelloRequest,::javaDemo::v1::HelloResponse>* server_unary_streamer) = 0;
  };
  typedef WithStreamedUnaryMethod_HelloWorld<Service > StreamedUnaryService;
  typedef Service SplitStreamedService;
  typedef WithStreamedUnaryMethod_HelloWorld<Service > StreamedService;
};

}  // namespace v1
}  // namespace javaDemo


#endif  // GRPC_javaDemo_2fv1_2fjavaDemo_2eproto__INCLUDED
